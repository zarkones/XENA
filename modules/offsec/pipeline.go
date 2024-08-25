package offsec

import (
	"common/debug"
	"common/machine"
	"common/sec"
	"common/sec/wordlists"
	"common/slices"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	sliceUtil "golang.org/x/exp/slices"

	c2api "github.com/zarkones/xena-client"
)

func runPipeline(pipeline c2api.Pipeline) (executedPipeline c2api.Pipeline) {
	var settings c2api.PipelineSettings
	if err := json.Unmarshal([]byte(pipeline.Settings), &settings); err != nil {
		return executedPipeline
	}

	defer debug.Println("Finished running the pipeline")
	debug.Println("Starting the pipeline")

	parentSteps := map[string]c2api.PipelineStep{}
	convergingSteps := map[string]c2api.PipelineStep{}
	allSteps := map[string]c2api.PipelineStep{}
	executedSteps := map[string]c2api.PipelineStep{}

	// Categorize steps, so we know which to execute in parallel.
	stepFrequencyMap := map[string]int{}
	for _, step := range settings.Steps {
		if len(step.LinkedTo) == 0 {
			continue
		}
		for _, linkedTo := range step.LinkedTo {
			stepFrequencyMap[linkedTo]++
		}
	}
	for _, step := range settings.Steps {
		allSteps[step.ID] = step
		freq := stepFrequencyMap[step.ID]
		if freq == 0 {
			parentSteps[step.ID] = step
			continue
		}
		if freq >= 2 {
			convergingSteps[step.ID] = step
			continue
		}
	}

	// parseInput gets outputs of other steps and pipeline variables and modifies the input based on that.
	parseInput := func(input, stepID string) string {
		// Replacing "__LINK" with output of connected step.
		if strings.Contains(input, "__LINK") {
			replacement := ""
			for _, executedStep := range executedSteps {
				if !sliceUtil.Contains(executedStep.LinkedTo, stepID) {
					continue
				}
				if stdout, ok := executedStep.Tool.Outputs["stdout"]; ok {
					replacement += stdout.Value + "\n"
				}
			}
			replacement = strings.TrimSuffix(replacement, "\n")
			input = strings.ReplaceAll(input, "__LINK", replacement)
		}
		for identifier, value := range settings.Input {
			input = strings.ReplaceAll(input, identifier, value)
		}
		return input
	}

	// TODO: Figure how to handle cnverging steps.
	var executeStep func(currStep c2api.PipelineStep)
	executeStep = func(currStep c2api.PipelineStep) {
		defer func() {
			executedSteps[currStep.ID] = currStep

			if len(currStep.LinkedTo) == 0 {
				return
			}

			// Execute linked steps in parallel.
			subStepsWg := sync.WaitGroup{}
			for _, linkedTo := range currStep.LinkedTo {
				subStepsWg.Add(1)
				go func(linkedTo string) {
					defer subStepsWg.Done()
					executeStep(allSteps[linkedTo])
				}(linkedTo)
			}
			subStepsWg.Wait()
		}()

		debug.Println("Running step:", currStep.Tool.Name)

		if currStep.Tool.ID == "" {
			println("herer")
		}

		currStep.Tool.Outputs = map[string]c2api.ToolOutput{}

		switch currStep.Tool.ID {

		default:
			currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: "not recognized"}
			return

		case "STR_CONTAINS":
			text := parseInput(currStep.Tool.Inputs["text"].Value, currStep.ID)
			subString := parseInput(currStep.Tool.Inputs["subString"].Value, currStep.ID)
			rawPerLine := parseInput(currStep.Tool.Inputs["perLine"].Value, currStep.ID)
			rawPerLine = strings.ToLower(rawPerLine)
			toTest := []string{text}
			switch rawPerLine {
			case "true", "1", "yes":
				toTest = strings.Split(text, "\n")
			}
			output := []string{}
			for _, line := range toTest {
				if strings.Contains(line, subString) {
					output = append(output, line)
				}
			}
			serializedOutput := strings.Join(output, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedOutput}
			return

		case "DOWNLOAD_FILE":
			url := parseInput(currStep.Tool.Inputs["url"].Value, currStep.ID)
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			c := http.Client{
				Timeout: time.Second * time.Duration(timeout),
			}
			resp, err := c.Get(url)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: string(respBody)}
			return

		case "SUBDOMAIN_ENUM_PASSIVE":
			rawDomains := parseInput(currStep.Tool.Inputs["domains"].Value, currStep.ID)
			domains := strings.Split(rawDomains, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			subdomains := []string{}
			analysis := []string{}
			errs := []error{}
			for _, domain := range domains {
				domain = strings.TrimPrefix(domain, "http://")
				domain = strings.TrimPrefix(domain, "https://")
				if domain == "" {
					continue
				}
				for _, page := range []int{1, 2, 3, 4} {
					urls, newAnalysis, err := sec.WebSearch("site:"+domain, page, time.Second*time.Duration(timeout), rpm)
					if err != nil {
						errs = append(errs, err)
						continue
					}
					for _, rawUrl := range urls {
						parsedUrl, err := url.Parse(rawUrl)
						if err != nil {
							errs = append(errs, err)
							continue
						}
						subdomains = append(subdomains, parsedUrl.Hostname())
					}
					analysis = append(analysis, newAnalysis...)
				}
			}
			subdomains = slices.Deduplicate(subdomains)
			serializedSubdomains := strings.Join(subdomains, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			serializedErr := ""
			if newErr := errors.Join(errs...); newErr != nil {
				serializedErr = newErr.Error()
			}
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedSubdomains}
			currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: serializedErr}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "WEB_SEARCH":
			query := parseInput(currStep.Tool.Inputs["query"].Value, currStep.ID)
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			errs := []error{}

			newPaths, newAnalysis, err := sec.WebSearch(query, 0, time.Second*time.Duration(timeout), rpm)
			if err != nil {
				errs = append(errs, err)
			}
			paths = append(paths, newPaths...)
			analysis = append(analysis, newAnalysis...)

			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			serializedErr := ""
			if newErr := errors.Join(errs...); newErr != nil {
				serializedErr = newErr.Error()
			}
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: serializedErr}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "SUBDOMAIN_ENUM_TOP100", "SUBDOMAIN_ENUM_TOP500", "SUBDOMAIN_ENUM_TOP1K", "SUBDOMAIN_ENUM_TOP10K":
			wordlist := func() *[]string {
				switch currStep.Tool.ID {
				default:
					return &wordlists.SubdomainsTop100
				case "SUBDOMAIN_ENUM_TOP500":
					return &wordlists.SubdomainsTop500
				case "SUBDOMAIN_ENUM_TOP1K":
					return &wordlists.SubdomainsTop1000
				case "SUBDOMAIN_ENUM_TOP10K":
					return &wordlists.SubdomainsTop10000
				}
			}()
			rawDomains := parseInput(currStep.Tool.Inputs["domains"].Value, currStep.ID)
			domains := strings.Split(rawDomains, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawThreads := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			threads, err := strconv.Atoi(rawThreads)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			subdomains := []string{}
			analysis := []string{}
			enumWG := sync.WaitGroup{}
			enumMux := sync.Mutex{}
			chunkedDomains := slices.Chunk(domains, threads)
			for _, chunk := range chunkedDomains {
				enumWG.Add(1)
				go func(chunk []string) {
					defer enumWG.Done()
					for _, domain := range chunk {
						domain = strings.TrimPrefix(domain, "http://")
						domain = strings.TrimPrefix(domain, "https://")
						if domain == "" {
							continue
						}
						newSubdomains, newAnalysis, _ := sec.SubdomainEnum(domain, wordlist, time.Second*time.Duration(timeout), rpm)
						enumMux.Lock()
						subdomains = append(subdomains, newSubdomains...)
						analysis = append(analysis, newAnalysis...)
						enumMux.Unlock()
					}
				}(chunk)
			}
			enumWG.Wait()
			serializedSubdomains := strings.Join(subdomains, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedSubdomains}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

			// This node is just an input, no need to process it.
		case "FILE":
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: currStep.Tool.Inputs["file"].Value}
			return

		case "GEN_IP_RANGE", "HTTP_REQUEST", "FTP_BRUTEFORCE", "SMB_BRUTEFORCE":
			// TODO
			currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: "not implemented"}
			return

		case "WEB_BYPASS_403":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			errs := []error{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, err := sec.Bypass403(url, time.Second*time.Duration(timeout), rpm)
				if err != nil {
					errs = append(errs, err)
				}
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			serializedErr := ""
			if newErr := errors.Join(errs...); newErr != nil {
				serializedErr = newErr.Error()
			}
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: serializedErr}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "HOST_HEADER_INJECTION_SCANNER":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			errs := []error{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, err := sec.HostHeaderInjection(url, time.Second*time.Duration(timeout), rpm)
				if err != nil {
					errs = append(errs, err)
				}
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			serializedErr := ""
			if newErr := errors.Join(errs...); newErr != nil {
				serializedErr = newErr.Error()
			}
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: serializedErr}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "CORS_SCANNER":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, _ := sec.AnalyzeCors(url, time.Second*time.Duration(timeout), rpm)
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "WEB_ENUM_FILE_UPLOAD":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, _ := sec.WebFindFileUpload(url, time.Second*time.Duration(timeout), rpm)
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "FIND_SECRETS":
			text := parseInput(currStep.Tool.Inputs["text"].Value, currStep.ID)
			secrets, err := sec.FindSecret(&text)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			serializedAnalysis := strings.Join(secrets, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "WEB_ENUM_GIT_DIR":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, _ := sec.WebFindGitDir(url, time.Second*time.Duration(timeout), rpm)
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "SSH_BRUTEFORCE":
			rawHosts := parseInput(currStep.Tool.Inputs["hosts"].Value, currStep.ID)
			hosts := strings.Split(rawHosts, "\n")
			port := parseInput(currStep.Tool.Inputs["port"].Value, currStep.ID)
			rawUsernames := parseInput(currStep.Tool.Inputs["usernameWordlist"].Value, currStep.ID)
			rawPasswords := parseInput(currStep.Tool.Inputs["passwordWordlist"].Value, currStep.ID)
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			usernames := strings.Split(rawUsernames, "\n")
			passwords := strings.Split(rawPasswords, "\n")
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			credentials := []string{}
			errs := []string{}
			for _, host := range hosts {
				if host == "" {
					continue
				}
				username, password, err := sec.SshBruteforce(host, port, &usernames, &passwords, time.Second*time.Duration(timeout), rpm)
				if err != nil {
					errs = append(errs, host+":"+err.Error())
					continue
				}
				credentials = append(credentials, username+" | "+password+" | "+net.JoinHostPort(host, port))
			}
			serializedErrors := strings.Join(errs, "\n")
			serializedCredentials := strings.Join(credentials, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedCredentials}
			currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: serializedErrors}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedCredentials}
			return

		case "WEB_ENUM_DEV_LEFTOVER":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, _ := sec.WebFindDevLeftover(url, time.Second*time.Duration(timeout), rpm)
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "WEB_ENUM_LOGS":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, _ := sec.WebFindLogs(url, time.Second*time.Duration(timeout), rpm)
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "WEB_ENUM_ADMIN_PANEL":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, _ := sec.WebFindAdminPanels(url, time.Second*time.Duration(timeout), rpm)
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "WEB_ENUM_TOP10K":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, _ := sec.WebFindTop10K(url, time.Second*time.Duration(timeout), rpm)
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "WEB_ENUM_BACKUPS":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, _ := sec.WebFindBackups(url, time.Second*time.Duration(timeout), rpm)
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "WEB_ENUM_CUSTOM":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			rawWordlist := parseInput(currStep.Tool.Inputs["wordlist"].Value, currStep.ID)
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			wordlist := strings.Split(rawWordlist, "\n")
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			paths := []string{}
			analysis := []string{}
			for _, url := range urls {
				if url == "" {
					continue
				}
				newPaths, newAnalysis, _ := sec.WebPathEnum(http.MethodGet, url, &wordlist, time.Second*time.Duration(timeout), rpm)
				paths = append(paths, newPaths...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedPaths := strings.Join(paths, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPaths}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "EXEC_SHELL":
			input := parseInput(currStep.Tool.Inputs["command"].Value, currStep.ID)
			output, err := machine.RunTerminal(input)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
			}
			output = strings.TrimSuffix(output, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: output}
			return

		case "DEDUPLICATE":
			input := parseInput(currStep.Tool.Inputs["text"].Value, currStep.ID)
			input = strings.TrimPrefix(
				strings.TrimSuffix(input, "\n"),
				"\n",
			)
			lines := strings.Split(input, "\n")
			uniqueLines := slices.Deduplicate(lines)
			serializedUniqueLines := strings.Join(uniqueLines, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedUniqueLines}
			return

		case "TRIM_SUFFIX":
			text := parseInput(currStep.Tool.Inputs["text"].Value, currStep.ID)
			suffix := parseInput(currStep.Tool.Inputs["suffix"].Value, currStep.ID)
			output := strings.TrimSuffix(text, suffix)
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: output}
			return

		case "TRIM_PREFIX":
			text := parseInput(currStep.Tool.Inputs["text"].Value, currStep.ID)
			prefix := parseInput(currStep.Tool.Inputs["prefix"].Value, currStep.ID)
			output := strings.TrimPrefix(text, prefix)
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: output}
			return

		case "READ_FILE":
			filePath := parseInput(currStep.Tool.Inputs["filePath"].Value, currStep.ID)
			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
			}
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: string(fileContent)}
			return

		case "WRITE_FILE":
			filePath := parseInput(currStep.Tool.Inputs["filePath"].Value, currStep.ID)
			fileContent := parseInput(currStep.Tool.Inputs["fileContent"].Value, currStep.ID)
			if err := os.WriteFile(filePath, []byte(fileContent), 0777); err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
			}
			return

		case "READ_DIR":
			dirPath := parseInput(currStep.Tool.Inputs["dirPath"].Value, currStep.ID)
			dirData, err := os.ReadDir(dirPath)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
			}
			dirEntries := ""
			for _, entry := range dirData {
				dirEntries += entry.Name() + "\n"
			}
			dirEntries = strings.TrimSuffix(dirEntries, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: dirEntries}
			return

		case "WALK_DIR":
			dirPath := parseInput(currStep.Tool.Inputs["dirPath"].Value, currStep.ID)
			filePaths, err := machine.WalkDir(dirPath)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
			}
			serializedFilePaths := strings.Join(filePaths, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedFilePaths}
			return

		// case "OS_SHUTDOWN":
		// 	if err := coldfire.Shutdown(); err != nil {
		// 		currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
		// 	}
		// 	return

		case "TLD_ENUM":
			rawNames := parseInput(currStep.Tool.Inputs["names"].Value, currStep.ID)
			names := strings.Split(rawNames, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawWordlist := parseInput(currStep.Tool.Inputs["wordlist"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			rawWordlist = strings.TrimSuffix(rawWordlist, "\n")
			wordlist := strings.Split(rawWordlist, "\n")
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			reachableDomains := []string{}
			analysis := []string{}
			for _, name := range names {
				if name == "" {
					continue
				}
				newDomains, newAnalysis, _ := sec.TldEnum(name, &wordlist, time.Second*time.Duration(timeout), rpm)
				reachableDomains = append(reachableDomains, newDomains...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedDomains := strings.Join(reachableDomains, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedDomains}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "SUBDOMAIN_ENUM":
			rawDomains := parseInput(currStep.Tool.Inputs["domains"].Value, currStep.ID)
			domains := strings.Split(rawDomains, "\n")
			rawTimeout := parseInput(currStep.Tool.Inputs["timeout"].Value, currStep.ID)
			rawWordlist := parseInput(currStep.Tool.Inputs["wordlist"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			rawWordlist = strings.TrimSuffix(rawWordlist, "\n")
			wordlist := strings.Split(rawWordlist, "\n")
			timeout, err := strconv.Atoi(rawTimeout)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			subdomains := []string{}
			analysis := []string{}
			for _, domain := range domains {
				domain = strings.TrimPrefix(domain, "http://")
				domain = strings.TrimPrefix(domain, "https://")
				if domain == "" {
					continue
				}
				newSubdomains, newAnalysis, _ := sec.SubdomainEnum(domain, &wordlist, time.Second*time.Duration(timeout), rpm)
				subdomains = append(subdomains, newSubdomains...)
				analysis = append(analysis, newAnalysis...)
			}
			serializedSubdomains := strings.Join(subdomains, "\n")
			serializedAnalysis := strings.Join(analysis, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedSubdomains}
			currStep.Tool.Outputs["analysis"] = c2api.ToolOutput{Type: "STRING", Value: serializedAnalysis}
			return

		case "WEB_CRAWL":
			rawUrls := parseInput(currStep.Tool.Inputs["urls"].Value, currStep.ID)
			urls := strings.Split(rawUrls, "\n")
			output := ""
			for _, url := range urls {
				if url == "" {
					continue
				}
				output += "\n" + webCrawl(url)
			}
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: output}
			return

		case "PORT_SCANNER":
			rawHosts := parseInput(currStep.Tool.Inputs["hosts"].Value, currStep.ID)
			hosts := strings.Split(rawHosts, "\n")
			rawPorts := parseInput(currStep.Tool.Inputs["ports"].Value, currStep.ID)
			rawRPM := parseInput(currStep.Tool.Inputs["rpm"].Value, currStep.ID)
			rpm, err := strconv.Atoi(rawRPM)
			if err != nil {
				currStep.Tool.Outputs["stderr"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
				return
			}
			rawPorts = strings.ReplaceAll(rawPorts, "\n", "")
			ports := strings.Split(rawPorts, ",")
			openPorts := []string{}
			sleepAmount := sec.GetSleep(rpm)
			for _, host := range hosts {
				if host == "" {
					continue
				}
				for _, rawPort := range ports {
					port, err := strconv.Atoi(rawPort)
					if err != nil {
						continue
					}
					sec.RateLimitedAction(sleepAmount, func() {
						if open := sec.PortScan(host, port); open {
							openPorts = append(openPorts, net.JoinHostPort(host, rawPort))
						}
					})
				}
			}
			serializedPorts := strings.Join(openPorts, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedPorts}
			return

		case "REVERSE_DNS":
			rawAddresses := parseInput(currStep.Tool.Inputs["addresses"].Value, currStep.ID)
			addresses := strings.Split(rawAddresses, "\n")
			domains := []string{}
			for _, address := range addresses {
				resolvedDomains, _ := sec.ReverseDNSLookup(address)
				domains = append(domains, resolvedDomains...)
			}
			serializedDomains := strings.Join(domains, "\n")
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: serializedDomains}
			return

		case "GATEWAY_IP_ADDR":
			address, err := sec.GetGatewayIP()
			if err != nil {
				currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: err.Error()}
			}
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: address}
			return

		case "GEN_RAND_IP":
			randomIP := strconv.Itoa(rand.Intn(255)) + "." + strconv.Itoa(rand.Intn(255)) + "." + strconv.Itoa(rand.Intn(255)) + "." + strconv.Itoa(rand.Intn(255))
			currStep.Tool.Outputs["stdout"] = c2api.ToolOutput{Type: "STRING", Value: randomIP}
			return

		}
	}

	wg := sync.WaitGroup{}

	for _, parentStep := range parentSteps {
		wg.Add(1)
		go func(step c2api.PipelineStep) {
			defer wg.Done()
			executeStep(step)
		}(parentStep)
	}

	wg.Wait()

	newSettings := c2api.PipelineSettings{
		Input: settings.Input,
		Steps: make(map[string]c2api.PipelineStep, len(executedSteps)),
	}
	for key, val := range executedSteps {
		newSettings.Steps[key] = val
	}

	jsonResponseSettings, err := json.Marshal(&newSettings)
	if err != nil {
		return executedPipeline
	}
	executedPipeline.Settings = string(jsonResponseSettings)
	return executedPipeline
}

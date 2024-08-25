package main

import (
	"common/builder"
	"common/machine"
	"common/slices"
	"log"
	"os"
	"tests/common"
)

func buildXENA() (err error) {
	if err := os.RemoveAll(common.PROJECT_EXPORT_PATH); err != nil {
		log.Println("project export path could not be removed:", err)
	}

	if err := os.MkdirAll(common.PROJECT_EXPORT_PATH, 0777); err != nil {
		log.Println("failed to create directory of the project export path")
		return err
	}

	if err := buildAgents(); err != nil {
		return err
	}

	if err := buildC2(); err != nil {
		return err
	}

	if err := buildUI(); err != nil {
		return err
	}

	return nil
}

func buildUI() (err error) {
	buildMap := map[string][]string{
		"linux": {"amd64", "386"},
	}

	for osName, architectures := range buildMap {
		log.Println("building", osName, "UI")

		for _, archName := range architectures {
			exportName := "XENA_" + osName + `_` + archName
			if osName == "windows" {
				exportName += ".exe"
			}
			buildCmd := `go build -ldflags="-s -w" -tags "netgo,` + osName + `" -o ` + common.UI_EXPORT_PATH + `/` + exportName + ` ./ui`

			env := slices.Deduplicate(append(os.Environ(), "CGO_ENABLED=1", "GOOS="+osName, "GOARCH="+archName))

			output, err := machine.RunTerminal(buildCmd, env...)
			if err == nil {
				continue
			}

			log.Println("build command ("+osName, archName+"):", buildCmd)
			log.Println("building output:\n", output)
			log.Fatalln("compilation of UI for", osName, archName, "has failed:", err)
		}
	}

	return nil
}

func buildC2() (err error) {
	buildMap := map[string][]string{
		"linux":   {"amd64", "386", "arm", "arm64"},
		"windows": {"amd64"},
		"darwin":  {"amd64", "arm64"},
	}

	for osName, architectures := range buildMap {
		log.Println("building", osName, "C2")

		for _, archName := range architectures {
			extraLDFlags := ""
			if osName == "windows" {
				extraLDFlags = "-H=windowsgui"
			}
			exportName := osName + `_` + archName
			if osName == "windows" {
				exportName += ".exe"
			}
			buildCmd := `go build -ldflags="-s -w -extldflags='-static' ` + extraLDFlags + `" -tags "netgo,` + osName + `,purego" -o ` + common.C2_EXPORT_PATH + `/` + exportName + ` ./c2`

			env := slices.Deduplicate(append(os.Environ(), "CGO_ENABLED=0", "GOOS="+osName, "GOARCH="+archName))

			output, err := machine.RunTerminal(buildCmd, env...)
			if err == nil {
				continue
			}

			log.Println("build command ("+osName, archName+"):", buildCmd)
			log.Println("building output:\n", output)
			log.Fatalln("compilation of C2 for", osName, archName, "has failed:", err)
		}
	}

	return nil
}

func buildAgents() (err error) {

	// Monolith agent.
	for osName, architectures := range builder.AgentBuildMap {
		log.Println("building", osName, "agents")

		for _, archName := range architectures {
			extraLDFlags := ""
			if osName == "windows" {
				extraLDFlags = "-H=windowsgui"
			}
			exportName := osName + `_` + archName
			if osName == "windows" {
				exportName += ".exe"
			}
			buildCmd := `go build -ldflags="-s -w -extldflags='-static' ` + extraLDFlags + `" -tags "netgo,` + osName + `,purego" -o ` + common.AGENT_EXPORT_PATH + `/` + exportName + ` ./agent`

			env := slices.Deduplicate(append(os.Environ(), "CGO_ENABLED=0", "GOOS="+osName, "GOARCH="+archName))

			output, err := machine.RunTerminal(buildCmd, env...)
			if err == nil {
				continue
			}

			log.Println("build command ("+osName, archName+"):", buildCmd)
			log.Println("building output:\n", output)
			log.Fatalln("compilation of agent for", osName, archName, "has failed:", err)
		}
	}

	// Modular agent.
	for osName, architectures := range builder.AgentBuildMap {
		if osName != "windows" && osName != "linux" {
			continue
		}

		log.Println("building modular", osName, "agents")

		for _, archName := range architectures {
			extraLDFlags := ""
			if osName == "windows" {
				extraLDFlags = "-H=windowsgui"
			}
			exportName := "modular_" + osName + `_` + archName
			if osName == "windows" {
				exportName += ".exe"
			}
			buildCmd := `go build -ldflags="-s -w -extldflags='-static' ` + extraLDFlags + `" -tags "netgo,xena_modular,` + osName + `,purego" -o ` + common.AGENT_EXPORT_PATH + `/` + exportName + ` ./agent`

			env := slices.Deduplicate(append(os.Environ(), "CGO_ENABLED=0", "GOOS="+osName, "GOARCH="+archName))

			output, err := machine.RunTerminal(buildCmd, env...)
			if err == nil {
				continue
			}

			log.Println("build command ("+osName, archName+"):", buildCmd)
			log.Println("building output:\n", output)
			log.Fatalln("compilation of agent for", osName, archName, "has failed:", err)
		}
	}

	return nil
}

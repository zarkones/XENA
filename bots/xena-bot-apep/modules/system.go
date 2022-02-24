package modules

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
)

// Structure describing basic system's information.
type OsDetails struct {
	Os       string `json:"os"`
	Arch     string `json:"arch"`
	CpuCount int    `json:"cpuCount"`
}

// GetOsDetails returns basic system's information.
func GetOsDetails() OsDetails {
	osName := runtime.GOOS
	if osName == "darwin" {
		osName = "mac"
	}

	arch := runtime.GOARCH

	cpuCount := runtime.NumCPU()

	return OsDetails{
		Os:       osName,
		Arch:     arch,
		CpuCount: cpuCount,
	}
}

// RunTerminal takes an input and runs it within the shell context and returns the result as a string.
func RunTerminal(input string) (string, error) {
	cmd := exec.Command(strings.TrimSuffix(input, "\n"))
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

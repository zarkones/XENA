package cases

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func AgentMainFileValidation() error {
	agentMainRaw, err := os.ReadFile("agent/main.go")
	if err != nil {
		return err
	}

	for i, line := range strings.Split(string(agentMainRaw), "\n") {
		lineN := strconv.Itoa(i + 1)

		if strings.Contains(line, "config.Patch") && strings.Contains(line, "//") {
			return errors.New("at line #" + lineN + " agents var patching has been disabled")
		}

		if strings.Contains(line, "PUBLIC KEY") && !strings.Contains(line, "//") {
			return errors.New("at line #" + lineN + " hardcoded key")
		}
	}

	return nil
}

package shell

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func Run(input string) (string, error) {
	if len(input) == 0 {
		return "<nil>", errors.New("IERMINAL_INPUT_INVALID")
	}

	cmd := exec.Command("bash", "-c", strings.TrimSuffix(input, "\n"))
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return out.String(), err
	}

	return strings.TrimSuffix(out.String(), "\n"), nil
}

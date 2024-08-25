package machine

import (
	"bytes"
	"errors"
	"os/exec"
	"runtime"
	"strings"
)

// RunTerminal takes an input and runs it within the shell context and returns the result as a string.
func RunTerminal(input string, env ...string) (output string, err error) {
	input = strings.TrimSuffix(input, "\n")

	cmd := func() *exec.Cmd {
		switch runtime.GOOS {
		case "windows":
			return exec.Command("cmd", "/C", input)
		default:
			return exec.Command("bash", "-c", input)
		}
	}()

	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr

	cmd.Env = env

	if currErr := cmd.Start(); currErr != nil {
		if err == nil {
			err = currErr
		} else {
			err = errors.Join(err, errors.New("cmd.Start: "), currErr)
		}
	}

	if currErr := cmd.Wait(); currErr != nil {
		if err == nil {
			err = currErr
		} else {
			err = errors.Join(err, errors.New("cmd.Wait: "), currErr)
		}
	}

	output = strings.Join([]string{out.String(), outErr.String()}, "\n")

	output = strings.TrimPrefix(output, "\n")
	output = strings.TrimSuffix(output, "\n")

	return output, err
}

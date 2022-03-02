package docker

import (
	"errors"
	"main/shell"
	"strings"
)

func CreateNetwork() error {
	if _, err := shell.Run("docker network create xena"); err != nil && !strings.HasPrefix(err.Error(), "exit status 1") {
		return errors.New("unable to create a docker network named 'xena'. Process exited with a status code different from 0 and 1")
	}
	return nil
}

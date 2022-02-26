package docker

import (
	"main/shell"
	"strings"
)

func CreateNetwork() error {
	if _, err := shell.Run("docker network create xena"); err != nil && !strings.HasPrefix(err.Error(), "exit status 1") {
		return err
	}
	return nil
}

package docker

import (
	"errors"
	"main/shell"
	"strings"
)

func Download() error {
	if version, err := shell.Run("docker -v"); err != nil || !strings.HasPrefix(version, "Docker version") {
		_, err = shell.Run("sudo apt install docker-ce docker-ce-cli containerd.io")
		if err != nil {
			return errors.New("unable to download docker-ce docker-ce-cli containerd.io using the apt package manager")
		}
	}
	return nil
}

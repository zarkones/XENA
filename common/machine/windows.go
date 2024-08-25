//go:build windows

package machine

import "golang.org/x/sys/windows/registry"

func UserLevelReg(path string, regKey string) error {
	key, err := registry.OpenKey(
		registry.CURRENT_USER,
		"Software\\Microsoft\\Windows\\CurrentVersion\\Run",
		registry.SET_VALUE|registry.ALL_ACCESS|registry.QUERY_VALUE,
	)
	if err != nil {
		return err
	}

	defer key.Close()

	if err := key.SetStringValue(regKey, path); err != nil {
		return err
	}

	return nil
}

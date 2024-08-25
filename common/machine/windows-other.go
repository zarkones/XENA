//go:build !windows

package machine

import "errors"

func UserLevelReg(path string, regKey string) (err error) {
	return errors.New("only on Windows")
}

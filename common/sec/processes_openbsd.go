//go:build openbsd
// +build openbsd

package sec

import "errors"

func Processes() (processes map[string][]int, err error) {
	return map[string][]int{}, errors.New("os not supported")
}

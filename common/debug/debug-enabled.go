//go:build xena_dbg
// +build xena_dbg

package debug

import "fmt"

func Println(a ...any) (n int, err error) {
	return fmt.Println(a...)
}

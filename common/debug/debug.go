//go:build !xena_dbg
// +build !xena_dbg

package debug

func Println(a ...any) (n int, err error) {
	return 0, nil
}

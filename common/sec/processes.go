//go:build !openbsd
// +build !openbsd

package sec

import "github.com/mitchellh/go-ps"

func Processes() (processes map[string][]int, err error) {
	ps, err := ps.Processes()
	if err != nil {
		return nil, err
	}
	processes = make(map[string][]int)
	for _, proc := range ps {
		procName := proc.Executable()
		pid := proc.Pid()
		if _, ok := processes[procName]; !ok {
			processes[procName] = []int{pid}
			continue
		}
		processes[procName] = append(processes[procName], pid)
	}
	return processes, nil
}

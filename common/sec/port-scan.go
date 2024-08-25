package sec

import (
	"time"

	portscanner "github.com/anvie/port-scanner"
)

func PortScan(address string, port int) bool {
	ps := portscanner.NewPortScanner(address, time.Second*10, 1)
	if ps == nil {
		return false
	}
	openPorts := ps.GetOpenedPort(port-1, port+1)
	return len(openPorts) != 0
}

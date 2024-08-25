package sec

import "net"

func ReverseDNSLookup(ipAddress string) (domains []string, err error) {
	ips, err := net.LookupAddr(ipAddress)
	if err != nil {
		return nil, err
	}
	return ips, nil
}

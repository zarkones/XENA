package cases

import (
	"common/machine"
	"log"
	"runtime"
	"strings"
	"tests/common"
	"time"

	"github.com/google/uuid"
	xenaC2 "github.com/zarkones/xena-client"
	cry "github.com/zarkones/xena-crypto"
)

func C2() error {
	var AUTH_TOKEN = "ABC123123i09iij9hiubh8ugyvgytgfttfrdd"
	var HOST = "127.0.0.1"
	var PORT = "8080"
	var BASE_URL = "http://" + HOST + ":" + PORT
	var GIN_MODE = "release"

	env := []string{
		"AUTH_TOKEN=" + AUTH_TOKEN,
		"HOST=" + HOST,
		"PORT=" + PORT,
		"GIN_MODE=" + GIN_MODE,
	}

	go func() {
		defer machine.RunTerminal("killall " + runtime.GOOS + "_" + runtime.GOARCH)

		if output, err := machine.RunTerminal(
			"."+common.SLASH+common.C2_EXPORT_PATH+common.SLASH+runtime.GOOS+"_"+runtime.GOARCH,
			env...,
		); err != nil {
			log.Println("error while running C2:", err)
			log.Println(output)
			log.Fatalln("critical failure of C2")
		}
	}()

	time.Sleep(time.Second * 10)

	c2PubKey, err := xenaC2.GetC2PublicKey()
	if err != nil {
		log.Println("failed to get C2 public key:", err)
		return err
	}

	if strings.Contains(string(c2PubKey), "PRIVATE") {
		log.Fatalln("C2 public key contains private key")
	}

	trustedPubKey, err := cry.ImportPubKeyPEM(c2PubKey)
	if err != nil {
		log.Println("C2 key nor parsable:", err)
		return err
	}

	xenaC2.AuthToken = &AUTH_TOKEN
	xenaC2.Init(&BASE_URL, trustedPubKey, time.Second*5)

	agentPrivKey, err := cry.GenPrivKey()
	if err != nil {
		log.Fatalln("agent's private key generation failed:", err)
	}
	if agentPrivKey == nil {
		log.Fatalln("agent's private key is null:", agentPrivKey)
	}

	agentPubKeyPEM, err := cry.PubKeyToPEM(&agentPrivKey.PublicKey)
	if err != nil {
		log.Fatalln("agent's public key serialization failed:", err)
	}

	const AGENT_HOSTNAME = "testacc"
	const AGENT_OS = "windows"
	const AGENT_ARCH = "amd64"

	agentID, err := xenaC2.Identify(AGENT_HOSTNAME, AGENT_OS, AGENT_ARCH, agentPubKeyPEM, agentPrivKey)
	if err != nil {
		log.Fatalln("agent identification failed:", err)
	}

	if len(agentID) <= 11 {
		log.Fatalln("agent ID seems invalid:", agentID)
	}

	if err := xenaC2.InsertMessage(xenaC2.Message{
		ID:      uuid.NewString(),
		AgentID: agentID,
		Request: "hostname",
	}); err != nil {
		log.Fatalln("failed to issue message as operator:", agentID)
	}

	messages, err := xenaC2.AgentFetchMessages(agentID, agentPrivKey)
	if err != nil {
		log.Fatalln("failed to fetch messages as an agent:", err)
	}

	if len(messages) != 1 {
		log.Fatalln("unexpected messages length as an agent:", len(messages))
	}

	if err := xenaC2.AgentRespondToMessage(messages[0].ID, "", AGENT_HOSTNAME); err != nil {
		log.Fatalln("failed to respond to a message as an agent:", err)
	}

	messages, err = xenaC2.FetchMessages(agentID)
	if err != nil {
		log.Fatalln("failed to fetch messages as an operator:", err)
	}

	if len(messages) != 1 {
		log.Fatalln("unexpected messages length as an operator:", len(messages))
	}

	if messages[0].AgentID != agentID {
		log.Fatalln("message's agent ID does not equal to agent's real ID:", messages[0].AgentID, "vs.", agentID)
	}
	if messages[0].Response != AGENT_HOSTNAME {
		log.Fatalln("unexpected message response:", messages[0].Response)
	}

	return nil
}

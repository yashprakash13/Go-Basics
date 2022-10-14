package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var Token string
var persistentMessageID string

func init() {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}

func main() {
	Token = os.Getenv("TOKEN")
	// Create a new Discord session using the provided bot token.
	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	bot.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	bot.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = bot.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()

}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.ChannelID == "1023789932772343868" {
		if persistentMessageID == "" {
			// read the last message id from the file and then delete it
			data, err_readfile := os.ReadFile("lastTextSentDetails.txt")
			if err_readfile != nil {
				fmt.Println("Error in reading file: ", err_readfile)
			} else {
				persistentMessageID = string(data)
			}
		}
		// delete the last message
		err := s.ChannelMessageDelete("1023789932772343868", persistentMessageID)
		if err != nil {
			fmt.Println("Erorr in deleting the previous sticky message: ", err)
		}
		// send a new message to the bottom
		messsage, err := s.ChannelMessageSend(m.ChannelID, "Here is a nice little sticky message.")

		// if message wasn't sent due to some reason
		if err != nil {
			fmt.Println("Erorr in sending sticky message: ", err)
		}

		// store the id of the new message sent in the persistent variable
		persistentMessageID = messsage.ID

		// write persistentMessageID string to file
		err = os.WriteFile("lastTextSentDetails.txt", []byte(persistentMessageID), 0644)
		if err != nil {
			fmt.Println("Error writing to file: ", err)
		}

	}

}

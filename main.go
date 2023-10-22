package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/quintui/songbird/internal/config"
)

func main() {

	config.Load()

	ds, err := connect(config.Get().AccessToken)

	if err != nil {
		log.Fatalf("Something went wrong connecting to Discord: %s", err)
	}

	ds.AddHandler(handleMessage)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	ds.Close()
}

func handleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	commandPrefix := config.Get().CommandPrefix
	messageContent := message.Content

	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(messageContent, commandPrefix) {
		command := strings.Replace(strings.Split(messageContent, " ")[0], commandPrefix, "", 1)
		switch command {
		case "play", "yt", "youtube":
			var query string
			if splittedMessage := strings.Split(messageContent, " "); len(splittedMessage) < 2 {
				session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("@%s Bimbo... It looks like you forgot to put a LINK.", message.Author.Username))
				return
			}
			query = strings.Replace(strings.Split(messageContent, " ")[1], commandPrefix, "", 1)

			handleYoutubeCommand(session, message, query)

		}

	}
}

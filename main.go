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

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	commandPrefix := config.Get().CommandPrefix
	message := m.Content

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(message, commandPrefix) {

		command := strings.Replace(strings.Split(message, " ")[0], commandPrefix, "", 1)

		switch command {
		case "play", "yt", "youtube":
			var query string
			fmt.Print(len(strings.Split(message, " ")))
			if splittedMessage := strings.Split(message, " "); len(splittedMessage) < 2 {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("@%s Bimbo... It looks like you forgot to put a LINK.", m.Author.Username))
				return
			}
			query = strings.Replace(strings.Split(message, " ")[1], commandPrefix, "", 1)

			handleYoutubeCommand(s, m, query)

		case "pong":
			s.ChannelMessageSend(m.ChannelID, "Ping")

		}

	}
}

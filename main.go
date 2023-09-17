package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/quintui/songbird/internal/config"
	"github.com/wader/goutubedl"
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
		query := strings.Replace(strings.Split(message, " ")[1], commandPrefix, "", 1)

		switch command {
		case "play", "yt", "youtube":
			handleYoutubeCommand(m, query)

		case "pong":
			s.ChannelMessageSend(m.ChannelID, "Ping")

		}

	}
}

func handleYoutubeCommand(session *discordgo.Session, message *discordgo.MessageCreate, query string) {
	session.ChannelMessageSend(message.ChannelID, "l")

	youtubeOptions := goutubedl.Options{}

	searchResult, err := goutubedl.New(context.Background(), query, youtubeOptions)

	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Oopsie @%s, Something went wrong with getting the video. URL: %s ", message.Author.Username, query))
		return
	}

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Hey, @%s, I found you video!!! %s", message.Author.Username))

	_, err = searchResult.Download(context.Background(), "")

	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Oopsie @%s, Something went wrong with downloading the video. URL: %s ", message.Author.Username, query))
		return
	}

}

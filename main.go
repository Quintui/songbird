package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
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

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "$ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

	if m.Content == "$pong" {
		s.ChannelMessageSend(m.ChannelID, "ping")
	}

}

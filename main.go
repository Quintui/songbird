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

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
		return
	}

	accessToken := os.Getenv("ACCESS_TOKEN")

	if accessToken == "" {
		log.Fatal("ACCESS_TOKEN not set")
		return
	}

	ds, err := discordgo.New("Bot " + accessToken)

	if err != nil {
		fmt.Println(err)
		return
	}

	ds.AddHandler(handleMessage)

	ds.Identify.Intents = discordgo.IntentsGuildMessages

	err = ds.Open()
	if err != nil {

		log.Fatal("Error opening connection", err)
		return
	}

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

package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func connect(accessToken string) *discordgo.Session {
	discordgo.MakeIntent(discordgo.IntentsGuildVoiceStates)
	ds, err := discordgo.New("Bot " + accessToken)

	if err != nil {
		log.Fatalf("Something went wrong connecting to Discord: %s", err)
	}

	err = ds.Open()

	if err != nil {
		log.Fatalf("Something went wrong connecting to Discord: %s", err)
	}

	return ds

}

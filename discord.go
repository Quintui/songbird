package main

import (
	"github.com/bwmarrin/discordgo"
)

func connect(accessToken string) (*discordgo.Session, error) {
	discordgo.MakeIntent(discordgo.IntentsGuildVoiceStates)
	ds, err := discordgo.New("Bot " + accessToken)

	if err != nil {
		return nil, err
	}

	err = ds.Open()

	return ds, err

}

package main

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

func getAuthorCurrentChannel(userId string) (*discordgo.Channel, error) {
	for _, guild := range dgSession.State.Guilds {
		for _, voice := range guild.VoiceStates {
			if voice.UserID == userId {
				channel, err := dgSession.State.Channel(voice.ChannelID)

				if err != nil {
					return nil, err
				}

				return channel, nil
			}
		}

	}

	return nil, errors.New("there is no user in voice channels ")
}

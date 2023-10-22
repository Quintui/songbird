package main

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

func getAuthorCurrentChannel(user *discordgo.User, guild *discordgo.Guild, session *discordgo.Session) (*discordgo.Channel, error) {
	for _, voiceConnection := range guild.VoiceStates {

		if voiceConnection.UserID == user.ID {
			channel, err := session.State.Channel(voiceConnection.ChannelID)

			if err != nil {
				return nil, err
			}

			return channel, nil

		}
	}

	return nil, errors.New("there is no user in voice channels ")
}

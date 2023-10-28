package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func playReporter(voiceInstance *VoiceInstance, message *discordgo.MessageCreate) {
	if voiceInstance == nil {
		dg.ChannelMessageSend(message.ChannelID, "I'm not in a voice channel!")
		return
	}

	authorChannel, err := SearchAuthorCurrentChannel(message.Author.ID)

	if err != nil {
		dg.ChannelMessageSend(message.ChannelID, err.Error())
		return
	}

	if voiceInstance.voice.ChannelID != authorChannel.ID {
		dg.ChannelMessageSend(message.ChannelID, "User is not connected to the channel")
		return
	}

	// send play my_song_youtube
	command := strings.SplitAfter(message.Content, strings.Fields(message.Content)[0])
	query := strings.TrimSpace(command[1])

	song, err := GetYoutubeVideo(query, voiceInstance, *message)

	if err != nil {
		dg.ChannelMessageSend(message.ChannelID, err.Error())
		return
	}

	go func() {
		songChan <- song
	}()

}

// Implement playing playlists

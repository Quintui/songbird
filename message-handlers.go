package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func playReporter(voiceInstance *VoiceInstance, message *discordgo.MessageCreate) {
	if voiceInstance == nil {
		// TODO: SEND ERROR
		return
	}
	if len(strings.Fields(message.Content)) < 2 {
		// TODO: SEND ERROR
		return
	}

	authorChannel, err := getAuthorCurrentChannel(message.Author.ID)

	if err != nil {
		// TODO: SEND ERROR
		return
	}

	// send play my_song_youtube
	command := strings.SplitAfter(message.Content, strings.Fields(message.Content)[0])
	query := strings.TrimSpace(command[1])

}

// USER CANNOT Play 2 songs in a row need to restart the bot
// Think about streaming songs instead of saving them to disk
// Implement queue
// Implement pause/resume
// Implement skip
// Implement playing playlists

package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

func handleYoutubeCommand(voiceInstance VoiceInstance, message *discordgo.MessageCreate) {
	messageAuthor := message.Author
	channel, err := session.State.Channel(message.ChannelID)
	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintln("Error with getting a message channel"))
		return
	}

	guild, err := session.State.Guild(channel.GuildID)

	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintln("Error with getting Guild"))
		return
	}

	authorCurrentChannel, err := getAuthorCurrentChannel(messageAuthor, guild, session)
	// USER CANNOT Play 2 songs in a row need to restart the bot
	// Think about streaming songs instead of saving them to disk
	// Implement queue
	// Implement pause/resume
	// Implement skip
	// Implement playing playlists

	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintln("Error: I don't know where to connect to  "))
		return
	}

	voiceConnection, err := session.ChannelVoiceJoin(authorCurrentChannel.GuildID, authorCurrentChannel.ID, false, true)

	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Listen @%s, Something went wrong with getting a channel, tell @%s that he is a bimbo and must to fix it...", message.Author.Username, "quinttt__"))
		return
	}

	downloadedSong, err := GetYoutubeVideo(query, voiceConnection, *message)

	voiceConnection.Speaking(true)

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Hey, @%s, I found your video!!! %s", message.Author.Username, query))

	done := make(chan error)
	dca.NewStream(source, vc, done)
	voiceConnection.Speaking(false)

	defer voiceConnection.Close()
}

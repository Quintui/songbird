package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func handleYoutubeCommand(session *discordgo.Session, message *discordgo.MessageCreate, query string) {

	// session.ChannelVoiceJoin()

	channel, err := session.State.Channel(message.ChannelID)

	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Listen @%s, Something went wrong with getting a channel, tell @%s that he is a bimbo and must to fix it...", message.Author.Username, "quinttt__"))
		return
	}

	// guild, err := session.State.Guild(channel.GuildID)

	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Listen @%s, Something went wrong with getting a guild, tell @%s that he is a bimbo and must to fix it...", message.Author.Username, "quinttt__"))
		return
	}

	voiceConnection, err := session.ChannelVoiceJoin(channel.GuildID, channel.ID, false, true)

	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Listen @%s, Something went wrong with connection to the channel, tell @%s that he is a bimbo and must to fix it...", message.Author.Username, "quinttt__"))
		return
	}

	defer voiceConnection.Close()

	// youtubeOptions := goutubedl.Options{}

	// searchResult, err := goutubedl.New(context.Background(), query, youtubeOptions)

	// if err != nil {
	// 	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Oopsie @%s, Something went wrong with getting the video. URL: %s ", message.Author.Username, query))
	// 	return
	// }

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Hey, @%s, I found you video!!! %s", message.Author.Username, query))
	//
	// downloadedVideo, err := searchResult.Download(context.Background(), "")

	// if err != nil {
	// 	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Oopsie @%s, Something went wrong with downloading the video. URL: %s ", message.Author.Username, query))
	// 	return
	// }

	// defer downloadedVideo.Close()
	//
}

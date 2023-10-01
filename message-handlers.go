package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/wader/goutubedl"
)

func handleYoutubeCommand(session *discordgo.Session, message *discordgo.MessageCreate, query string) {
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

	currentChannel, err := getCurrentChannel(messageAuthor, guild, session)

	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintln("Error: I don't know where to connect to  "))
		return
	}

	voiceConnection, err := session.ChannelVoiceJoin(currentChannel.GuildID, currentChannel.ID, false, true)

	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Listen @%s, Something went wrong with getting a channel, tell @%s that he is a bimbo and must to fix it...", message.Author.Username, "quinttt__"))
		return
	}

	defer voiceConnection.Close()

	youtubeOptions := goutubedl.Options{}

	searchResult, err := goutubedl.New(context.Background(), query, youtubeOptions)

	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Oopsie @%s, Something went wrong with getting the video. URL: %s ", message.Author.Username, query))
		return
	}

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Hey, @%s, I found you video!!! %s", message.Author.Username, query))
	//
	downloadedVideo, err := searchResult.Download(context.Background(), "")

	// downloadedVideo.Read()

	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Oopsie @%s, Something went wrong with downloading the video. URL: %s ", message.Author.Username, query))
		return
	}

	defer downloadedVideo.Close()
	//
}

func getCurrentChannel(user *discordgo.User, guild *discordgo.Guild, session *discordgo.Session) (*discordgo.Channel, error) {
	for _, voiceConnection := range guild.VoiceStates {

		if voiceConnection.UserID == user.ID {
			channel, err := session.State.Channel(voiceConnection.ChannelID)

			if err != nil {
				return nil, err
			}

			return channel, nil

		}
		return nil, errors.New("There is no user in voice channels ")

	}

	return nil, nil
}

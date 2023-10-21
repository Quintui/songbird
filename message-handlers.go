package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/kkdai/youtube/v2"
)

type youtubeSong struct {
	URL string
}

func (ys *youtubeSong) downloadVideo() (string, error) {

	client := youtube.Client{}
	parsedURL, err := url.Parse(ys.URL)
	if err != nil {
		log.Printf("Unable to decode url: %s, err: %v", ys.URL, err)
		return "", err
	}

	videoID := parsedURL.Query().Get("v")
	youtubeVideo, err := client.GetVideo(videoID)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	formats := youtubeVideo.Formats.WithAudioChannels()

	stream, _, err := client.GetStream(youtubeVideo, &formats[0])
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fileName := fmt.Sprintf("%s.mp4", videoID)

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, stream)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return fileName, nil
}

func handleYoutubeCommand(session *discordgo.Session, message *discordgo.MessageCreate, query string) {
	messageAuthor := message.Author
	channel, err := session.State.Channel(message.ChannelID)
	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintln("Error with getting a message channel"))
		return
	}

	youtubeSong := youtubeSong{
		URL: query,
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

	downloadedSong, err := youtubeSong.downloadVideo()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer os.Remove(downloadedSong)

	voiceConnection, err := session.ChannelVoiceJoin(authorCurrentChannel.GuildID, authorCurrentChannel.ID, false, true)

	if err != nil {
		fmt.Println(err)
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Listen @%s, Something went wrong with getting a channel, tell @%s that he is a bimbo and must to fix it...", message.Author.Username, "quinttt__"))
		return
	}

	voiceConnection.Speaking(true)

	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Hey, @%s, I found your video!!! %s", message.Author.Username, query))

	dgvoice.PlayAudioFile(voiceConnection, downloadedSong, make(<-chan bool))

	voiceConnection.Speaking(false)

	defer voiceConnection.Close()
}

package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/bwmarrin/discordgo"
	"github.com/kkdai/youtube/v2"
)

func GetYoutubeVideo(youtubeUrl string, voiceInstance *VoiceInstance, message discordgo.MessageCreate) (songStruct PkgSong, err error) {

	client := youtube.Client{}
	parsedURL, err := url.Parse(youtubeUrl)
	if err != nil {
		log.Printf("Unable to decode url: %s, err: %v", youtubeUrl, err)
		return
	}

	videoID := parsedURL.Query().Get("v")
	youtubeVideo, err := client.GetVideo(videoID)

	if err != nil {
		fmt.Println(err)
		return
	}

	formats := youtubeVideo.Formats.WithAudioChannels()

	downloadedVideoUrl, err := client.GetStreamURL(youtubeVideo, &formats[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	song := Song{
		Title:       youtubeVideo.Title,
		Description: youtubeVideo.Description,
		ChannelId:   message.ChannelID,
		VideoUrl:    downloadedVideoUrl,
	}

	songStruct.v = voiceInstance
	songStruct.song = song

	return
}

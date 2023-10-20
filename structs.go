package main

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

type Song struct {
	Title       string
	Description string
	ChannelId   string
	VideoUrl    string
}

type VoiceInstance struct {
	voice      *discordgo.VoiceConnection
	session    *discordgo.Session
	queueMutex *sync.Mutex
	songMutex  *sync.Mutex
	channelID  string
	guildId    string
	nowPlaying Song
	pause      bool
	queue      []Song
	speaking   bool
}

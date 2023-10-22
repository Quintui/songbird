package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"sync"
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
	encoder    *dca.EncodeSession
	stream     *dca.StreamingSession
	queueMutex *sync.Mutex
	songMutex  *sync.Mutex
	channelID  string
	guildId    string
	nowPlaying Song
	pause      bool
	queue      []Song
	speaking   bool
}

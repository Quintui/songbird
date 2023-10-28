package main

import (
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
)

type Song struct {
	Title       string
	Description string
	ChannelId   string
	VideoUrl    string
}

type PkgSong struct {
	song Song
	v    *VoiceInstance
}

type VoiceInstance struct {
	voice      *discordgo.VoiceConnection
	session    *discordgo.Session
	encoder    *dca.EncodeSession
	stream     *dca.StreamingSession
	queueMutex sync.Mutex
	songMutex  sync.Mutex
	// channelID  string
	guildId    string
	nowPlaying Song
	stop       bool
	skip       bool
	pause      bool
	queue      []Song
	speaking   bool
}

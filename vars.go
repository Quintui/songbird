package main

import "github.com/bwmarrin/discordgo"

var (
	dgSession      *discordgo.Session
	voiceInstances = map[string]*VoiceInstance{}
	songChan       = make(chan PkgSong)
	// mutex          sync.Mutex
)

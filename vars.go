package main

import (
	"sync"

	"github.com/bwmarrin/discordgo"
)

var (
	dg             *discordgo.Session
	voiceInstances = map[string]*VoiceInstance{}
	songChan       = make(chan PkgSong)
	mutex          sync.Mutex
)

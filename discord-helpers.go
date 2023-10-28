package main

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
)

func initRoutine() {
	songChan = make(chan PkgSong)
	go GlobalPlay(songChan)
}

func SearchAuthorCurrentChannel(userId string) (*discordgo.Channel, error) {
	for _, guild := range dg.State.Guilds {
		for _, voice := range guild.VoiceStates {
			if voice.UserID == userId {
				channel, err := dg.State.Channel(voice.ChannelID)

				if err != nil {
					return nil, err
				}

				return channel, nil
			}
		}

	}

	return nil, errors.New("there is no user in voice channels ")
}

func SearchGuildId(textChannelId string) string {
	channel, _ := dg.Channel(textChannelId)

	return channel.GuildID
}

func CreateVoiceInstance(session *discordgo.Session, message *discordgo.MessageCreate) {
	guildID := SearchGuildId(message.ChannelID)
	authorVoiceChannel, err := SearchAuthorCurrentChannel(message.Author.ID)

	if err != nil {
		log.Println("ERROR: You need to join a voice channel! ", err)
		return
	}
	// create new voice instance
	mutex.Lock()
	defer mutex.Unlock()
	createdVoiceInstance := new(VoiceInstance)
	createdVoiceInstance.guildId = guildID
	createdVoiceInstance.session = session
	createdVoiceInstance.voice, err = dg.ChannelVoiceJoin(guildID, authorVoiceChannel.ID, false, true)

	if err != nil {
		log.Println("ERROR: Error to join in a voice channel: ", err)
		return

	}

	voiceInstances[guildID] = createdVoiceInstance
}

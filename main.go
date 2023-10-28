package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/quintui/songbird/internal/config"
)

func main() {

	config.Load()

	dg = connect(config.Get().AccessToken)

	dg.AddHandler(handleMessage)

	initRoutine()
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
}

func handleMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	commandPrefix := config.Get().CommandPrefix
	messageContent := message.Content

	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(messageContent, commandPrefix) {

		command := strings.Replace(strings.Split(messageContent, " ")[0], commandPrefix, "", 1)

		guildId := SearchGuildId(message.ChannelID)

		var voiceInstance *VoiceInstance

		voiceInstance = voiceInstances[guildId]

		if voiceInstance == nil {
			CreateVoiceInstance(session, message)
		}

		voiceInstance = voiceInstances[guildId]

		switch command {
		case "play", "yt", "youtube":
			playReporter(voiceInstance, message)
		}

	}
}

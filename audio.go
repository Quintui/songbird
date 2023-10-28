package main

import (
	"io"
	"log"

	"github.com/jonas747/dca"
)

func GlobalPlay(songChan chan PkgSong) {
	for {
		select {
		case song := <-songChan:
			go song.v.PlayQueue(song.song)
		}
	}

}

func (v *VoiceInstance) PlayQueue(song Song) {
	v.AddQueue(song)

	if v.speaking {
		return
	}

	go func() {

		v.songMutex.Lock()
		defer v.songMutex.Unlock()

		for {

			if len(v.queue) == 0 {
				return
			}

			v.nowPlaying = v.GetQueueSong()

			v.pause = false
			v.skip = false
			v.speaking = true
			v.voice.Speaking(true)

			v.DCA(v.nowPlaying.VideoUrl)
			v.QueueRemoveFirst()

			if v.stop {
				v.QueueClean()
			}
			v.stop = false
			v.skip = false
			v.speaking = false
			v.voice.Speaking(false)

		}

	}()

}

func (v *VoiceInstance) DCA(url string) {
	opts := dca.StdEncodeOptions
	opts.RawOutput = true
	opts.Bitrate = 96
	opts.Application = "lowdelay"

	encodeSession, err := dca.EncodeFile(url, opts)
	if err != nil {
		log.Println("FATA: Failed creating an encoding session: ", err)
	}
	v.encoder = encodeSession
	done := make(chan error)

	stream := dca.NewStream(encodeSession, v.voice, done)
	v.stream = stream

	for {
		select {
		case err := <-done:
			if err != nil || err != io.EOF {
				// FIX: ERROR HERE
				log.Println("FATA: An error occured", err)
			}
			encodeSession.Cleanup()
			return
		}
	}

}

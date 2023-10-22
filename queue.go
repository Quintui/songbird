package main

func (v *VoiceInstance) GetQueueSong() (song Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 {
		return v.queue[0]
	}
	return
}

func (v *VoiceInstance) AddQueue(song Song) {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()

	v.queue = append(v.queue, song)
}

func (v *VoiceInstance) QueueRemoveFirst() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	if len(v.queue) != 0 {
		v.queue = v.queue[1:]
	}
}

func (v *VoiceInstance) QueueClean() {
	v.queueMutex.Lock()
	defer v.queueMutex.Unlock()
	v.queue = []Song{}
}

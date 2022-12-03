package main

import (
	hook "github.com/robotn/gohook"
	"magro"
)

var recordCommands = []string{"ctrl", "alt", "r"}

func main() {
	eventCh := hook.Start()
	recorder := magro.NewRecorder(eventCh)
	defer recorder.Close()

	// Toggle isRecording (CTRL + ALT + R).
	hook.Register(hook.KeyDown, recordCommands, func(event hook.Event) {
		recorder.Toggle()
	})

	<-hook.Process(recorder.Start())
}

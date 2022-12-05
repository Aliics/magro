package main

import (
	hook "github.com/robotn/gohook"
	"magro"
	"magro/fyne"
)

var recordCommands = []string{"ctrl", "alt", "r"}

func main() {
	recorder := createRecorder()
	defer recorder.Close()

	window, refresh := fyne.CreateMainWindow(recorder)
	defer window.Close()

	go func() {
		// Refresh the window if recording state changes.

		knownRecording := recorder.IsRecording
		for {
			if recorder.IsRecording != knownRecording {
				knownRecording = recorder.IsRecording
				refresh()
			}
		}
	}()

	go func() {
		<-hook.Process(recorder.Start())
	}()

	window.ShowAndRun()
}

func createRecorder() *magro.Recorder {
	eventCh := hook.Start()
	recorder := magro.NewRecorder(eventCh)

	// Toggle isRecording (CTRL + ALT + R).
	hook.Register(hook.KeyDown, recordCommands, func(event hook.Event) {
		recorder.Toggle()
	})

	return recorder
}

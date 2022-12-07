package main

import (
	hook "github.com/robotn/gohook"
	"magro"
	"magro/gui"
)

func main() {
	recorder := createRecorder()
	defer recorder.Close()

	go func() {
		<-hook.Process(recorder.Start())
	}()

	mainGUI := gui.NewGUI(recorder)
	defer mainGUI.Close()

	mainGUI.ShowAndRun()
}

func createRecorder() *magro.Recorder {
	eventCh := hook.Start()
	recorder := magro.NewRecorder(eventCh)

	return recorder
}

package main

import (
	hook "github.com/robotn/gohook"
	"log"
	"magro"
	"magro/gui"
	"magro/persist"
)

func main() {
	persister, err := persist.NewPersister()
	if err != nil {
		log.Fatalln(err)
	}
	defer persister.Close()

	err = persister.Load()
	if err != nil {
		log.Fatalln(err)
	}

	recorder := createRecorder()
	defer recorder.Close()

	recorder.RecordedMacros = persister.RecordedMacros

	mainGUI := gui.NewGUI(recorder)
	defer mainGUI.Close()

	go func() {
		err = persister.StartAndPersist()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		<-hook.Process(recorder.Start())
	}()

	mainGUI.ShowAndRun()
}

func createRecorder() *magro.Recorder {
	eventCh := hook.Start()
	recorder := magro.NewRecorder(eventCh)

	return recorder
}

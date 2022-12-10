package main

import (
	hook "github.com/robotn/gohook"
	"log"
	"magro"
	"magro/gui"
	"magro/persist"
)

func main() {
	// Macros slices for the Persister and Recorder to share.
	var recordedMacros []magro.Macro

	persister, err := createPersister(&recordedMacros)
	if err != nil {
		log.Fatalln(err)
	}
	defer persister.Close()

	recorder := createRecorder()
	defer recorder.Close()
	recorder.RecordedMacros = &recordedMacros

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

func createPersister(recordedMacros *[]magro.Macro) (*persist.Persister, error) {
	persister, err := persist.NewPersister()
	if err != nil {
		return nil, err
	}
	persister.RecordedMacros = recordedMacros

	err = persister.Load()
	if err != nil {
		return nil, err
	}

	return persister, nil
}

func createRecorder() *magro.Recorder {
	eventCh := hook.Start()
	recorder := magro.NewRecorder(eventCh)

	return recorder
}

package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	hook "github.com/robotn/gohook"
	"magro"
)

var recordCommands = []string{"ctrl", "alt", "r"}

func main() {
	recorder := createRecorder()
	defer recorder.Close()

	go func() {
		<-hook.Process(recorder.Start())
	}()

	createAndRunWindow(recorder)
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

func createAndRunWindow(recorder *magro.Recorder) {
	window := app.New().NewWindow("magro")
	defer window.Close()

	macroList := widget.NewList(
		func() int {
			return len(recorder.RecordedMacros)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			object.(*widget.Label).SetText(fmt.Sprintf("Macro %d", id))
		},
	)

	var recordButton *widget.Button
	recordButton = widget.NewButton("Record", func() {
		recorder.Toggle()

		if recorder.IsRecording {
			recordButton.SetText("Record")
		} else {
			recordButton.SetText("Stop")
		}
	})

	window.SetContent(
		container.NewVBox(
			container.NewHBox(recordButton),
			macroList,
		),
	)

	window.ShowAndRun()
}

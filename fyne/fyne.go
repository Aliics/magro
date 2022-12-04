package fyne

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"magro"
)

func CreateWindow(recorder *magro.Recorder) fyne.Window {
	window := app.New().NewWindow("magro")

	macroList := widget.NewList(
		func() int {
			return len(recorder.RecordedMacros)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("template", nil)
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			button := object.(*widget.Button)
			button.SetText(fmt.Sprintf("Macro %d", id))
			button.OnTapped = func() {
				go playMacro(recorder.RecordedMacros[id])
			}
		},
	)

	var recordButton *widget.Button
	recordButton = widget.NewButton("Record (ctrl + alt + r)", func() {
		recorder.Toggle()

		if recorder.IsRecording {
			recordButton.SetText("Record (ctrl + alt + r)")
		} else {
			recordButton.SetText("Stop (ctrl + alt + r)")
		}
	})

	window.SetContent(
		container.NewBorder(
			container.NewVBox(
				recordButton,
				widget.NewSeparator(),
			),
			nil, nil, nil,
			macroList,
		),
	)

	return window
}

func playMacro(macro magro.Macro) {
	err := macro.PlayEvents()
	if err != nil {
		log.Fatalln(err)
	}
}

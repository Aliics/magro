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

func CreateWindow(recorder *magro.Recorder) (fyne.Window, func()) {
	window := app.New().NewWindow("magro")

	macroList := createMacroList(recorder)

	recordButton := widget.NewButton("Record (ctrl + alt + r)", func() {
		recorder.Toggle()
	})

	refreshContent := func() {
		if recorder.IsRecording {
			recordButton.SetText("Stop (ctrl + alt + r)")
		} else {
			recordButton.SetText("Record (ctrl + alt + r)")
		}

		window.Content().Refresh()
	}

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

	return window, refreshContent
}

func createMacroList(recorder *magro.Recorder) *widget.List {
	return widget.NewList(
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
}

func playMacro(macro magro.Macro) {
	err := macro.PlayEvents()
	if err != nil {
		log.Fatalln(err)
	}
}

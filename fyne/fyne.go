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

	recordButton := widget.NewButton("Record (ctrl + alt + r)", func() {
		recorder.Toggle()
	})
	setRecordButtonActive(recordButton, false)

	refreshContent := func() {
		setRecordButtonActive(recordButton, recorder.IsRecording)

		window.Content().Refresh()
	}

	macroList := createMacroList(recorder)

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

func setRecordButtonActive(button *widget.Button, active bool) {
	if active {
		button.SetText("Stop (ctrl + alt + r)")
		button.Importance = widget.HighImportance
	} else {
		button.SetText("Record (ctrl + alt + r)")
		button.Importance = widget.MediumImportance
	}
}

func createMacroList(recorder *magro.Recorder) *widget.List {
	return widget.NewList(
		func() int {
			return len(recorder.RecordedMacros)
		},
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil, nil,
				widget.NewLabel(""),
				widget.NewButton("Play", nil),
			)
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			macroItem := object.(*fyne.Container)

			label := macroItem.Objects[0].(*widget.Label)
			label.SetText(fmt.Sprintf("Macro %d", id))

			button := macroItem.Objects[1].(*widget.Button)
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

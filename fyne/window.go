package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"magro"
)

const windowTitle = "magro"

func CreateWindow(recorder *magro.Recorder) (fyne.Window, func()) {
	window := app.New().NewWindow(windowTitle)

	recordButton := createRecordButton(recorder)

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

	refreshContent := func() {
		setRecordButtonActive(recordButton, recorder.IsRecording)

		window.Content().Refresh()
	}

	return window, refreshContent
}

func playMacro(macro magro.Macro) {
	err := macro.PlayEvents()
	if err != nil {
		log.Fatalln(err)
	}
}

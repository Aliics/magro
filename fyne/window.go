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

func CreateMainWindow(recorder *magro.Recorder) (window fyne.Window, refreshMacroList func()) {
	window = app.New().NewWindow(windowTitle)
	window.Resize(fyne.NewSize(240, 360))

	// Main content.
	recordButton := createRecordButton(recorder)
	macroList := createMacroList(recorder)
	mainContent := container.NewBorder(
		container.NewVBox(
			recordButton,
			widget.NewSeparator(),
		),
		nil, nil, nil,
		macroList,
	)

	go handleContentSwitch(recorder, window, mainContent)

	// Show the main content.
	contentSwitchCh <- contentSwitch{kind: contentSwitchKindMain}

	refreshMacroList = func() {
		setRecordButtonActive(recordButton, recorder.IsRecording)

		window.Content().Refresh()
	}

	return window, refreshMacroList
}

func playMacro(macro *magro.Macro) {
	err := macro.PlayEvents()
	if err != nil {
		log.Fatalln(err)
	}
}

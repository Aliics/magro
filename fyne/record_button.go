package fyne

import (
	"fyne.io/fyne/v2/widget"
	"magro"
)

const (
	recordLabel = "Record (ctrl + alt + r)"
	stopLabel   = "Stop (ctrl + alt + r)"
)

func createRecordButton(recorder *magro.Recorder) *widget.Button {
	recordButton := widget.NewButton(recordLabel, func() {
		recorder.Toggle()
	})
	setRecordButtonActive(recordButton, false)

	return recordButton
}

func setRecordButtonActive(button *widget.Button, active bool) {
	if active {
		button.SetText(stopLabel)
		button.Importance = widget.HighImportance
	} else {
		button.SetText(recordLabel)
		button.Importance = widget.MediumImportance
	}
}

package fyne

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"magro"
)

func createMacroList(recorder *magro.Recorder) *widget.List {
	return widget.NewList(
		func() int {
			return len(recorder.RecordedMacros)
		},
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil, nil,
				newMacroLabel(""),
				widget.NewButton("Play", nil),
			)
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			macroItem := object.(*fyne.Container)

			label := macroItem.Objects[0].(*macroLabel)
			label.SetText(fmt.Sprintf("Macro %d", id))
			label.onDoubleTapped = func() {
				log.Printf("opening macro %d", id)
			}

			playButton := macroItem.Objects[1].(*widget.Button)
			playButton.OnTapped = func() {
				go playMacro(recorder.RecordedMacros[id])
			}
		},
	)
}

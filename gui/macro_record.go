package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
	"log"
)

var (
	startRecordIcon = theme.MediaRecordIcon()
	stopRecordIcon  = theme.MediaStopIcon()
)

func (g *GUI) createMacroRecord() *fyne.Container {
	macroList := widget.NewList(
		func() int {
			return len(*g.recorder.RecordedMacros)
		},
		func() fyne.CanvasObject {
			label := newTappableLabel("")
			button := widget.NewButtonWithIcon("", theme.MediaPlayIcon(), nil)
			return container.NewBorder(
				nil, nil,
				label, button,
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			macro := &(*g.recorder.RecordedMacros)[i]
			item := o.(*fyne.Container)

			label := item.Objects[0].(*tappableLabel)
			label.OnDoubleTapped = func() {
				details := newMacroDetails(macro, g.mainWindow, g.switchToMacroRecord, func() {
					*g.recorder.RecordedMacros = slices.Delete(*g.recorder.RecordedMacros, i, i+1)
					g.switchToMacroRecord()
				})
				g.mainWindow.SetContent(details.Container)
			}

			button := item.Objects[1].(*widget.Button)

			label.SetText(macro.Name)
			button.OnTapped = func() {
				err := macro.PlayEvents()
				if err != nil {
					log.Fatalln(err)
				}
			}
		},
	)

	recordButton := widget.NewButtonWithIcon("", startRecordIcon, nil)
	recordButton.OnTapped = func() {
		g.recorder.ToggleBlocking()

		// Switch icon for start/stop.
		if g.recorder.IsRecording {
			recordButton.Icon = stopRecordIcon
			recordButton.Importance = widget.HighImportance
		} else {
			recordButton.Icon = startRecordIcon
			recordButton.Importance = widget.MediumImportance
		}

		recordButton.Refresh()
		macroList.Refresh()
	}

	return container.NewBorder(
		recordButton,
		nil, nil, nil,
		macroList,
	)
}

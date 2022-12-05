package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"magro"
)

func createMacroContent(macro *magro.Macro) *fyne.Container {
	backButton := widget.NewButton("Back", func() {
		contentSwitchCh <- contentSwitch{kind: contentSwitchKindMain}
	})

	nameBinding := binding.BindString(&macro.Name)
	nameEntry := widget.NewEntryWithData(nameBinding)

	eventList := widget.NewList(
		func() int {
			return len(macro.Events)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			label := object.(*widget.Label)
			label.SetText(macro.Events[id].String())
		},
	)

	return container.NewBorder(
		container.NewVBox(backButton, nameEntry),
		nil, nil, nil,
		eventList,
	)
}

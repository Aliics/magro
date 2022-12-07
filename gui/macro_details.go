package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"magro"
)

type macroDetails struct {
	Container *fyne.Container

	macro               *magro.Macro
	switchToMacroRecord func()
}

func newMacroDetails(macro *magro.Macro, switchToMacroRecord func()) *macroDetails {
	m := &macroDetails{
		macro:               macro,
		switchToMacroRecord: switchToMacroRecord,
	}

	m.initContainer()

	return m
}

func (m *macroDetails) initContainer() {
	backButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		m.switchToMacroRecord()
	})
	nameEntry := widget.NewEntryWithData(binding.BindString(&m.macro.Name))
	topBorder := container.NewBorder(nil, nil, backButton, nil, nameEntry)

	eventList := widget.NewList(
		func() int {
			return len(m.macro.Events)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(m.macro.Events[i].String())
		},
	)

	m.Container = container.NewBorder(
		topBorder,
		nil, nil, nil,
		eventList,
	)
}

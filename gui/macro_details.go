package gui

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"magro"
	"strings"
)

var (
	errEmptyEntry = errors.New("text cannot be empty")
)

type macroDetails struct {
	Container *fyne.Container

	macro               *magro.Macro
	parentWindow        fyne.Window
	switchToMacroRecord func()
	deleteMacro         func()
}

func newMacroDetails(
	macro *magro.Macro,
	parentWindow fyne.Window,
	switchToMacroRecord, deleteMacro func(),
) *macroDetails {
	m := &macroDetails{
		macro:               macro,
		parentWindow:        parentWindow,
		switchToMacroRecord: switchToMacroRecord,
		deleteMacro:         deleteMacro,
	}

	m.initContainer()

	return m
}

func (m *macroDetails) initContainer() {
	backButton := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), m.switchToMacroRecord)

	nameEntry := widget.NewEntryWithData(binding.BindString(&m.macro.Name))
	nameEntry.Validator = func(name string) error {
		if strings.TrimSpace(name) == "" {
			backButton.Disable()
			return errEmptyEntry
		}

		backButton.Enable()
		return nil
	}

	deleteButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), m.showDeleteDialog)

	topBorder := container.NewBorder(nil, nil, backButton, deleteButton, nameEntry)

	eventList := m.createEventList()

	m.Container = container.NewBorder(
		topBorder,
		nil, nil, nil,
		eventList,
	)
}

func (m *macroDetails) showDeleteDialog() {
	dialog.
		NewConfirm(
			"Delete",
			fmt.Sprintf(`Delete "%s"?`, m.macro.Name),
			func(shouldDelete bool) {
				if shouldDelete {
					m.deleteMacro()
				}
			},
			m.parentWindow,
		).
		Show()
}

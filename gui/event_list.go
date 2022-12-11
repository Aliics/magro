package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/slices"
)

func (m *macroDetails) createEventList() *widget.List {
	return widget.NewList(
		func() int {
			return len(m.macro.Events)
		},
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil, nil,
				container.NewHBox(
					widget.NewLabel(""),
					widget.NewLabel(""),
				),
				container.NewHBox(
					widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil),
					widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
				),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			event := &m.macro.Events[i]
			item := o.(*fyne.Container)

			// Labels with Event data.
			{
				box := item.Objects[0].(*fyne.Container)

				deltaLabel := box.Objects[0].(*widget.Label)
				deltaLabel.SetText(fmt.Sprintf("%.3f", event.Delta.Seconds()))

				keyLabel := box.Objects[1].(*widget.Label)
				keyLabel.SetText(fmt.Sprintf("%s %c", event.KeyKind, event.Keycode))
			}

			// Actions container with edit and delete buttons.
			{
				box := item.Objects[1].(*fyne.Container)

				editButton := box.Objects[0].(*widget.Button)
				editButton.OnTapped = m.showEditMacro(event)

				deleteButton := box.Objects[1].(*widget.Button)
				deleteButton.OnTapped = func() {
					m.macro.Events = slices.Delete(m.macro.Events, i, i+1)
				}
			}
		},
	)
}

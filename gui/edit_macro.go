package gui

import (
	"errors"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"magro"
	"time"
)

var (
	errSingleCharExpected = errors.New("single character expected")
)

func (m *macroDetails) showEditMacro(event *magro.Event) func() {
	return func() {
		// Delta modification
		deltaEntry := widget.NewEntry()
		deltaEntry.SetText(event.Delta.String())
		deltaEntry.Validator = func(s string) error {
			_, err := time.ParseDuration(s)
			if err != nil {
				return err
			}

			return nil
		}
		deltaEntry.OnChanged = func(s string) {
			delta, _ := time.ParseDuration(s)
			event.Delta = delta
			m.parentWindow.Content().Refresh()
		}

		// KeyKind modification
		keyKindSelect := widget.NewSelect([]string{
			magro.KeyKindDown.String(),
			magro.KeyKindUp.String(),
		}, nil)
		keyKindSelect.Selected = event.KeyKind.String()
		keyKindSelect.OnChanged = func(s string) {
			keyKind, err := magro.NewKeyKindFromString(s)
			if err != nil {
				return
			}

			event.KeyKind = keyKind
			m.parentWindow.Content().Refresh()
		}

		// Keycode modification
		keycodeEntry := widget.NewEntry()
		keycodeEntry.SetText(string(event.Keycode))
		keycodeEntry.Validator = func(s string) error {
			if len(s) != 1 {
				return errSingleCharExpected
			}

			return nil
		}
		keycodeEntry.OnChanged = func(s string) {
			if err := keycodeEntry.Validate(); err != nil {
				return
			}

			event.Keycode = rune(s[0])
			m.parentWindow.Content().Refresh()
		}

		dialog.
			NewForm(
				"Edit",
				"Ok",
				"Cancel",
				[]*widget.FormItem{
					widget.NewFormItem("Delta", deltaEntry),
					widget.NewFormItem("Kind", keyKindSelect),
					widget.NewFormItem("Keycode", keycodeEntry),
				},
				nil,
				m.parentWindow,
			).
			Show()
	}
}

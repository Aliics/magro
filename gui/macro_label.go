package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type macroLabel struct {
	widget.Label

	OnDoubleTapped func()
}

func newMacroLabel(text string) *macroLabel {
	label := &macroLabel{}

	label.ExtendBaseWidget(label)
	label.SetText(text)

	return label
}

func (m *macroLabel) DoubleTapped(_ *fyne.PointEvent) {
	if m.OnDoubleTapped == nil {
		return
	}

	m.OnDoubleTapped()
}

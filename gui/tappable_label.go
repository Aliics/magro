package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type tappableLabel struct {
	widget.Label

	OnDoubleTapped func()
}

func newTappableLabel(text string) *tappableLabel {
	label := &tappableLabel{}

	label.ExtendBaseWidget(label)
	label.SetText(text)

	return label
}

func (l *tappableLabel) DoubleTapped(_ *fyne.PointEvent) {
	if l.OnDoubleTapped == nil {
		return
	}

	l.OnDoubleTapped()
}

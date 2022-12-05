package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"time"
)

const (
	doubleTapDebounce = 330 * time.Millisecond
)

type macroLabel struct {
	widget.Label

	onDoubleTapped func()
	lastTap        *time.Time
}

func newMacroLabel(text string) *macroLabel {
	t := &macroLabel{}
	t.Text = text

	t.ExtendBaseWidget(t)

	return t
}

func (b *macroLabel) Tapped(_ *fyne.PointEvent) {
	if b.onDoubleTapped != nil {
		if b.lastTap == nil || time.Now().Sub(*b.lastTap) > doubleTapDebounce {
			now := time.Now()
			b.lastTap = &now
		} else {
			b.lastTap = nil
			b.onDoubleTapped()
		}
	}
}

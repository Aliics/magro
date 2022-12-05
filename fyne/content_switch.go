package fyne

import (
	"fyne.io/fyne/v2"
	"magro"
)

var contentSwitchCh = make(chan contentSwitch)

type contentSwitch struct {
	kind       contentSwitchKind
	macroIndex int
}

type contentSwitchKind uint8

const (
	contentSwitchKindMain contentSwitchKind = iota
	contentSwitchKindMacro
)

func handleContentSwitch(recorder *magro.Recorder, window fyne.Window, mainContent *fyne.Container) {
	for content := range contentSwitchCh {
		switch content.kind {
		case contentSwitchKindMain:
			window.SetContent(mainContent)
		case contentSwitchKindMacro:
			window.SetContent(createMacroContent(recorder.RecordedMacros[content.macroIndex]))
		}
	}
}

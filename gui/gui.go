package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"magro"
)

// Initialize first to make sure fyne.App is ready.
var fyneApp = app.New()

var defaultWindowSize = fyne.NewSize(360, 480)

type GUI struct {
	mainWindow  fyne.Window
	macroRecord *fyne.Container

	recorder *magro.Recorder
}

func NewGUI(recorder *magro.Recorder) *GUI {
	mainWindow := fyneApp.NewWindow("magro")
	mainWindow.Resize(defaultWindowSize)

	gui := &GUI{
		mainWindow: mainWindow,
		recorder:   recorder,
	}

	gui.initWidgets()

	return gui
}

func (g *GUI) ShowAndRun() {
	g.mainWindow.ShowAndRun()
}

func (g *GUI) Close() {
	fyneApp.Quit()
}

func (g *GUI) switchToMacroRecord() {
	g.mainWindow.SetContent(g.macroRecord)
}

func (g *GUI) initWidgets() {
	g.macroRecord = g.createMacroRecord()
	g.switchToMacroRecord()
}

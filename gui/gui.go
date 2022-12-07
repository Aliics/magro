package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"magro"
)

// Initialize first to make sure fyne.App is ready.
var fyneApp = app.New()

var defaultWindowSize = fyne.NewSize(240, 360)

type GUI struct {
	mainWindow fyne.Window
	appTabs    *container.AppTabs

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

func (g *GUI) initWidgets() {
	listTab := container.NewTabItem("list", g.createMacroRecordList())

	g.appTabs = container.NewAppTabs(listTab)
	g.appTabs.SetTabLocation(container.TabLocationBottom)

	g.mainWindow.SetContent(g.appTabs)
}

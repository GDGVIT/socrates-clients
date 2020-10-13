package tui

import (
	"github.com/rivo/tview"
	"settings/tui/screen"
)

func Start() {
	app := tview.NewApplication()
	s := screen.New(app)
	app.SetInputCapture(s.HandleInput)
	s.Start()
}
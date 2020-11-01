package main

import (
	"log"
	"os"
	"github.com/rivo/tview"
	"settings/tui/screen"
)

func main() {
	// Setup logging
	file, err := os.OpenFile("socrates-client-logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.SetOutput(file)

	// Start application
	app := tview.NewApplication()
	s := screen.New(app)
	app.SetInputCapture(s.HandleInput)
	s.Start()
}
package main

import (
	"github.com/rainu/go-calendar/internal/ui"
	"time"
)

func main() {
	ParseFlags()

	app := ui.NewMainFrame(time.Now())

	app.ShowAndRun()
}

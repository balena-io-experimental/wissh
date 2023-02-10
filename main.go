package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Wissh")

	gui, err := NewGUI()
	if err != nil {
		// TODO: We should at least try to show this (also) on a Window.
		fmt.Fprintf(os.Stderr, "Oopsie, error initializing Wissh: %v\n", err)
		os.Exit(1)
	}
	gui.SetButtonAction(runChecksFunc(gui))

	w.SetContent(gui.Root)
	w.Resize(fyne.NewSize(1024.0, 700.0))
	w.ShowAndRun()
}

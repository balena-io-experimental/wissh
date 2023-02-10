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

	// w.SetContent(container.NewVBox(widget.NewLabel("É nóissh!")))
	gui, err := NewGUI()
	if err != nil {
		// TODO: We should at least try to show this (also) on a Window.
		fmt.Fprintf(os.Stderr, "Oopsie, error initializing Wissh: %v\n", err)
		os.Exit(1)
	}
	w.SetContent(gui.Root)
	w.Resize(fyne.NewSize(1024.0, 700.0))
	w.ShowAndRun()

	runner, err := NewSSHRunner("root", "192.168.100.80:22222")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Sorry, wissh failed to establish an ssh connection: %v", err)
		os.Exit(1)
	}

	cmd := "curl https://api.balena-cloud.com/ping"
	stdout, stderr, err := runner.Run(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd)
	fmt.Printf("--- stdout ---\n%v\n---stderr---\n%v\n------------\n", stdout, stderr)

	cmd = `nc -w 5 -G 1 cloudlink.balena-cloud.com 443 && echo "Reachable." || echo "Not reachable."`
	stdout, stderr, err = runner.Run(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd)
	fmt.Printf("--- stdout ---\n%v\n---stderr---\n%v\n------------\n", stdout, stderr)

	cmd = "curl -v https://registry2.balena-cloud.com"
	stdout, stderr, err = runner.Run(cmd)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd)
	fmt.Printf("--- stdout ---\n%v\n---stderr---\n%v\n------------\n", stdout, stderr)
}

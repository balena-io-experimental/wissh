package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type WisshGUI struct {
	Root fyne.CanvasObject
}

func NewGUI() *WisshGUI {
	gui := &WisshGUI{
		Root: container.NewVBox(
			widget.NewLabel("Settings"),
			newConfigSection(),
			widget.NewLabel("Results"),
			newResultsSection(),
		),
	}
	return gui
}

func newConfigSection() fyne.CanvasObject {
	deviceIPEntry := widget.NewEntry()
	deviceIPEntry.SetText("192.168.100.80")

	// TODO: Can users change the SSH port on a device? Maybe this should simply
	// be a constant. Or be hidden under some "advanced settings" things.
	sshPortEntry := widget.NewEntry()
	sshPortEntry.SetText("22222")

	sshKeyFileEntry := widget.NewEntry()
	sshKeyFileEntry.SetText("/home/lmb/.ssh/id_rsa_wissh_test")

	form := widget.NewForm(
		widget.NewFormItem("Device Local IP Address", deviceIPEntry),
		widget.NewFormItem("SSH Port", sshPortEntry),
		widget.NewFormItem("SSH Private Key File", sshKeyFileEntry),
	)

	return form
}

func newResultsSection() fyne.CanvasObject {
	results := widget.NewMultiLineEntry()
	results.Disable()
	return results
}

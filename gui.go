package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type WisshGUI struct {
	Root fyne.CanvasObject

	deviceIP   binding.String
	sshPort    binding.String
	sshKeyFile binding.String

	theButton  *widget.Button
	theResults *widget.Entry
}

func NewGUI() (*WisshGUI, error) {
	gui := &WisshGUI{}

	gui.deviceIP = binding.NewString()
	err := gui.deviceIP.Set("192.168.100.80")
	if err != nil {
		return nil, err
	}

	gui.sshPort = binding.NewString()
	err = gui.sshPort.Set("22222")
	if err != nil {
		return nil, err
	}

	gui.sshKeyFile = binding.NewString()

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	err = gui.sshKeyFile.Set(homeDir + "/.ssh/id_rsa")
	if err != nil {
		return nil, err
	}

	top := container.NewVBox()
	top.Add(widget.NewLabel("Configuration"))
	top.Add(newConfigSection(gui))
	top.Add(widget.NewLabel("Actions"))
	top.Add(newActionsSection(gui))
	top.Add(widget.NewLabel("Results"))

	results := newResultsSection(gui)
	root := container.NewBorder(top, nil, nil, nil, results)
	gui.Root = root

	return gui, nil
}

func (gui *WisshGUI) DeviceIP() string {
	ip, err := gui.deviceIP.Get()
	if err != nil {
		// TODO: Can't fail in the current implementation of Get(), but we
		// should handle this somehow better!
		return ""
	} else {
		return ip
	}
}

func (gui *WisshGUI) SSHPort() string {
	ip, err := gui.sshPort.Get()
	if err != nil {
		// TODO: Can't fail in the current implementation of Get(), but we
		// should handle this somehow better!
		return ""
	} else {
		return ip
	}
}

func (gui *WisshGUI) SSHKeyFile() string {
	ip, err := gui.sshKeyFile.Get()
	if err != nil {
		// TODO: Can't fail in the current implementation of Get(), but we
		// should handle this somehow better!
		return ""
	} else {
		return ip
	}
}

func (gui *WisshGUI) SetButtonAction(action func()) {
	gui.theButton.OnTapped = action
}

func newConfigSection(gui *WisshGUI) fyne.CanvasObject {
	deviceIPEntry := widget.NewEntry()
	deviceIPEntry.Bind(gui.deviceIP)

	// TODO: Can users change the SSH port on a device? Maybe this should simply
	// be a constant. Or be hidden under some "advanced settings" things.
	sshPortEntry := widget.NewEntry()
	sshPortEntry.Bind(gui.sshPort)

	sshKeyFileEntry := widget.NewEntry()
	sshKeyFileEntry.Bind(gui.sshKeyFile)

	form := widget.NewForm(
		widget.NewFormItem("Device Local IP Address", deviceIPEntry),
		widget.NewFormItem("SSH Port", sshPortEntry),
		widget.NewFormItem("SSH Private Key File", sshKeyFileEntry),
	)

	return form
}

func newActionsSection(gui *WisshGUI) fyne.CanvasObject {
	button := widget.NewButton("Diagnose!", func() {})
	gui.theButton = button
	return button
}

func newResultsSection(gui *WisshGUI) fyne.CanvasObject {
	results := widget.NewMultiLineEntry()
	gui.theResults = results
	return results
}

package main

import (
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
}

func NewGUI() (*WisshGUI, error) {
	root := container.NewVBox()
	gui := &WisshGUI{
		Root: root,
	}

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
	err = gui.sshKeyFile.Set("/home/lmb/.ssh/id_rsa_wissh_test")
	if err != nil {
		return nil, err
	}

	root.Add(widget.NewLabel("Configuration"))
	root.Add(newConfigSection(gui))
	root.Add(widget.NewLabel("Results"))
	root.Add(newResultsSection())

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

func newResultsSection() fyne.CanvasObject {
	results := widget.NewMultiLineEntry()
	results.Disable()
	return results
}

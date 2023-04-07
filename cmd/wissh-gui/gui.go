package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/balena-io-experimental/wissh/pkg/wissh"
)

type WisshGUI struct {
	Root fyne.CanvasObject

	deviceIP   binding.String
	sshPort    binding.String
	sshKeyFile binding.String

	theButton  *widget.Button
	theResults *fyne.Container
	mainWindow fyne.Window
}

func NewGUI(mainWindow fyne.Window) (*WisshGUI, error) {
	gui := &WisshGUI{
		mainWindow: mainWindow,
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

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	err = gui.sshKeyFile.Set(homeDir + "/.ssh/id_rsa")
	if err != nil {
		return nil, err
	}

	vbox := container.NewVBox()
	vbox.Add(widget.NewRichTextFromMarkdown("# Configuration"))
	vbox.Add(newConfigSection(gui))
	vbox.Add(widget.NewRichTextFromMarkdown("# Actions"))
	vbox.Add(newActionsSection(gui))
	vbox.Add(widget.NewRichTextFromMarkdown("# Results"))
	vbox.Add(newResultsSection(gui))

	top := container.NewVScroll(vbox)
	gui.Root = top

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
	top := container.NewHBox()
	button := widget.NewButton("Diagnose!", func() {})
	gui.theButton = button
	top.Add(button)
	return top
}

func newResultsSection(gui *WisshGUI) fyne.CanvasObject {
	results := container.NewVBox()
	gui.theResults = results
	return results
}

//
// Check UI
//

func newCheckUI(check wissh.Check, err error) fyne.CanvasObject {
	top := container.NewVBox()
	top.Add(widget.NewRichTextFromMarkdown("## " + check.Name()))
	status := "Passed!"

	if err != nil {
		status = fmt.Sprintf("Couldn't run the test: %v", err)
	} else if !check.Passed() {
		status = "FAILED!"
	}
	top.Add(widget.NewLabel(status))

	// TODO: Control flow is ugly here!
	if err != nil {
		return top
	}

	if ok, remarks := check.IlluminatingRemarks(); ok {
		mdView := widget.NewRichTextFromMarkdown(remarks)
		mdView.Wrapping = fyne.TextWrapWord
		top.Add(mdView)
	}
	if ok, details := check.Details(); ok {
		detailsBox := container.NewVBox()
		scroll := container.NewHScroll(widget.NewRichTextFromMarkdown(details))
		scroll.Hide()
		toggleButton := widget.NewButton("Show Details", nil)
		toggleButton.OnTapped = func() {
			if scroll.Visible() {
				scroll.Hide()
				toggleButton.SetText("Show Details")
			} else {
				scroll.Show()
				toggleButton.SetText("Hide Details")
			}
		}

		detailsBox.Add(container.NewHBox(toggleButton))
		detailsBox.Add(scroll)
		top.Add(detailsBox)
	}

	return top
}

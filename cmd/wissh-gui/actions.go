package main

import (
	"fmt"

	"fyne.io/fyne/v2/dialog"
	"github.com/balena-io-experimental/wissh/pkg/checks"
	"github.com/balena-io-experimental/wissh/pkg/wissh"
)

func runChecksFunc(gui *WisshGUI) func() {
	return func() {
		// TODO: Consider running these tests asynchronously, and make the UI
		// behave nicely as the checks run. For example, some kind of progress
		// reporting, and disabling the button during the process.

		deviceIP := gui.DeviceIP()
		sshPort := gui.SSHPort()
		sshKeyFile := gui.SSHKeyFile()

		gui.theResults.RemoveAll()

		oldText := gui.theButton.Text
		gui.theButton.SetText("Running...")
		gui.theButton.Disable()
		defer func() {
			gui.theButton.SetText(oldText)
			gui.theButton.Enable()
		}()

		if err := canSSHToDevice(deviceIP, sshPort, sshKeyFile); err != nil {
			dialog.ShowError(err, gui.mainWindow)
			return
		}

		for _, check := range checks.All(deviceIP, sshPort, sshKeyFile) {
			err := check.Run()
			gui.theResults.Add(newCheckUI(check, err))
		}
	}
}

// canSSHToDevice checks if we can SSH to to the device. If this fails, it
// doesn't make sense to even try running the actual checks (which all depend on
// SSH). The error return value explains what when wrong (likely to include not
// very user-friendly information). A nil return value means that we can SSH to
// the device.
func canSSHToDevice(deviceIP, sshPort, sshKeyFile string) error {
	runner, err := wissh.NewSSHRunner("root", deviceIP+":"+sshPort, sshKeyFile)
	if err != nil {
		return fmt.Errorf("Error while preparing to run SSH command: %w", err)
	}
	defer runner.Destroy()

	_, _, err = runner.Run("ls")
	if err != nil {
		return fmt.Errorf("Error while running SSH command: %w", err)
	}

	return nil
}

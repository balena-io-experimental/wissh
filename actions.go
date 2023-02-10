package main

import (
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func runChecksFunc(gui *WisshGUI) func() {
	return func() {
		gui.theResults.SetText("")

		printResult := func(s string) {
			gui.theResults.SetText(gui.theResults.Text + "\n" + s)
		}

		showError := func(msg string, err error) {
			dialog.ShowError(fmt.Errorf(msg+": %w", err), fyne.CurrentApp().Driver().AllWindows()[0])
		}

		runner, err := NewSSHRunner("root", gui.DeviceIP()+":"+gui.SSHPort(), gui.SSHKeyFile())
		if err != nil {
			showError("Failed to establish an ssh connection", err)
			return
		}

		runCommand := func(cmd string) error {
			stdout, stderr, err := runner.Run(cmd)
			if err != nil {
				showError("Failed to run an ssh command", err)
				return errors.New("failed to run command")
			}
			printResult(cmd)
			printResult(fmt.Sprintf("--- stdout ---\n%v\n---stderr---\n%v\n------------\n", stdout, stderr))
			return nil
		}

		runCommand(`curl https://api.balena-cloud.com/ping`)
		runCommand(`nc -w 5 -G 1 cloudlink.balena-cloud.com 443 && echo "Reachable." || echo "Not reachable."`)
		runCommand(`curl -v https://registry2.balena-cloud.com`)
	}
}

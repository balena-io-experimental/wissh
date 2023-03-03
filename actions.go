package main

func runChecksFunc(gui *WisshGUI) func() {
	return func() {
		// TODO: Before running any checks, try running a sanity check (like a
		// dummy ssh command), to catch any obvious errors that would cause all
		// tests to fail.

		deviceIP := gui.DeviceIP()
		sshPort := gui.SSHPort()
		sshKeyFile := gui.SSHKeyFile()

		gui.theResults.RemoveAll()

		check := newPingAPI(deviceIP, sshPort, sshKeyFile)
		err := check.Run()
		gui.theResults.Add(newCheckUI(check, err))
	}
}

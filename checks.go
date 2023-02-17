package main

type pingAPI struct {
	SSHCommand
}

func (c *pingAPI) Passed() bool {
	return c.StdOut == "OK"
}

func (c *pingAPI) IlluminatingRemarks() (bool, string) {
	if c.Passed() {
		return true,
			"We reached the balena API server.\n\n" +
				"This means the network path from the device to balenaCloud is working.\n"
	}

	return true,
		"We failed to reach the balenaAPI server.\n\n" +
			"This means there's something wrong on the path from the device to balenaCloud.\n" +
			"Perhaps you a have a firewall blocking outgoing requests to `https://api.balena-cloud.com/ping`?\n"
}

package checks

import "github.com/balena-io-experimental/wissh/pkg/wissh"

type pingAPI struct {
	wissh.SSHCommand
}

// Returns a check that tries to ping the balenaCloud API server.
func NewPingAPI(ip, port, sshKeyFile string) wissh.Check {
	return &pingAPI{
		SSHCommand: wissh.SSHCommand{
			Command:    "curl https://api.balena-cloud.com/ping",
			IP:         ip,
			Port:       port,
			SSHKeyFile: sshKeyFile,
		},
	}
}

func (c *pingAPI) Name() string {
	return "Connect With the API Server"
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
		"We failed to reach the balena API server.\n\n" +
			"This means there's something wrong on the path from the device to balenaCloud.\n" +
			"Perhaps you a have a firewall blocking outgoing requests to `https://api.balena-cloud.com/ping`?\n"
}

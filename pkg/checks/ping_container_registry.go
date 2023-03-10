package checks

import "github.com/balena-io-experimental/wissh/pkg/wissh"

type pingContainerRegistry struct {
	wissh.SSHCommand
}

// Returns a check that tries to ping the balena container registry.
func NewPingContainerRegistry(ip, port, sshKeyFile string) wissh.Check {
	return &pingContainerRegistry{
		SSHCommand: wissh.SSHCommand{
			Command:    "curl https://registry2.balena-cloud.com",
			IP:         ip,
			Port:       port,
			SSHKeyFile: sshKeyFile,
		},
	}
}

func (c *pingContainerRegistry) Name() string {
	return "Connect With the Container Registry"
}

func (c *pingContainerRegistry) Passed() bool {
	return c.ExitStatus == 0
}

func (c *pingContainerRegistry) IlluminatingRemarks() (bool, string) {
	if c.Passed() {
		return true,
			"We reached the balena container registry.\n\n" +
				"This means the device should be able to pull Docker images.\n"
	}

	return true,
		"We failed to reach the balena container container registry.\n\n" +
			"This means the device won't be able to pull Docker images.\n" +
			"Perhaps you a have a firewall blocking outgoing requests to `https://registry2.balena-cloud.com`?\n"
}

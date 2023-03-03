package main

//
// Ping API Server
//

type pingAPI struct {
	SSHCommand
}

func newPingAPI(ip, port, sshKeyFile string) *pingAPI {
	return &pingAPI{
		SSHCommand: SSHCommand{
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

//
// Ping Container Registry
//

type pingContainerRegistry struct {
	SSHCommand
}

func newPingContainerRegistry(ip, port, sshKeyFile string) *pingContainerRegistry {
	return &pingContainerRegistry{
		SSHCommand: SSHCommand{
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

// TODO: nc -w 5 -G 1 cloudlink.balena-cloud.com 443 && echo "Reachable." || echo "Not reachable."`)
// Or an equivalent that works...

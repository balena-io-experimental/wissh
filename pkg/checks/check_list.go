package checks

import "github.com/balena-io-experimental/wissh/pkg/wissh"

// All returns a slice with an instance of each check.
func All(deviceIP, sshPort, sshKeyFile string) []wissh.Check {
	return []wissh.Check{
		NewPingAPI(deviceIP, sshPort, sshKeyFile),
		NewPingContainerRegistry(deviceIP, sshPort, sshKeyFile),
	}
}

package checks

import (
	_ "embed"

	"github.com/balena-io-experimental/wissh/pkg/wissh"
)

//go:embed ping_container_registry_remarks_success.md
var pingContainerRegistryRemarksSuccess string

//go:embed ping_container_registry_remarks_failure.md
var pingContainerRegistryRemarksFailure string

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
		return true, pingContainerRegistryRemarksSuccess
	}

	return true, pingContainerRegistryRemarksFailure
}

package checks

import (
	_ "embed"

	"github.com/balena-io-experimental/wissh/pkg/wissh"
)

//go:embed reach_cloudlink_remarks_success.md
var reachCloudlinkRemarksSuccess string

//go:embed reach_cloudlink_remarks_failure.md
var reachCloudlinkRemarksFailure string

type reachCloudlink struct {
	wissh.SSHCommand
}

// Returns a check that tries to reach balena's cloudlink (AKA VPN).
func NewReachCLoudlink(ip, port, sshKeyFile string) wissh.Check {
	return &reachCloudlink{
		SSHCommand: wissh.SSHCommand{
			Command:    "echo | nc cloudlink.balena-cloud.com 443",
			IP:         ip,
			Port:       port,
			SSHKeyFile: sshKeyFile,
		},
	}
}

func (c *reachCloudlink) Name() string {
	return "Connect With balena's Cloudlink (AKA VPN)"
}

func (c *reachCloudlink) Passed() bool {
	return c.ExitStatus == 0
}

func (c *reachCloudlink) IlluminatingRemarks() (bool, string) {
	if c.Passed() {
		return true, reachCloudlinkRemarksSuccess
	}

	return true, reachCloudlinkRemarksFailure
}

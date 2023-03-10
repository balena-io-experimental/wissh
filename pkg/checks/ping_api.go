package checks

import (
	_ "embed"

	"github.com/balena-io-experimental/wissh/pkg/wissh"
)

//go:embed ping_api_remarks_success.md
var pingAPIRemarksSuccess string

//go:embed ping_api_remarks_failure.md
var pingAPIRemarksFailure string

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
		return true, pingAPIRemarksSuccess
	}

	return true, pingAPIRemarksFailure
}

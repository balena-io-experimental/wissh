package checks

import (
	"bufio"
	_ "embed"
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/balena-io-experimental/wissh/pkg/wissh"
	"gopkg.in/ini.v1"
)

//go:embed static_ip_config_remarks_success.md
var staticIPConfigRemarksSuccess string

//go:embed static_ip_config_remarks_failure.md
var staticIPConfigRemarksFailure string

//go:embed static_ip_config_remarks_no_static_found.md
var staticIPConfigRemarksNoStaticFound string

type staticIPConfig struct {
	wissh.SSHCommand
	passed            bool
	foundAnyStatic    bool
	additionalRemarks string
}

// Returns a check that looks for common errors in static IP configurations.
func NewStaticIPConfig(ip, port, sshKeyFile string) wissh.Check {
	return &staticIPConfig{
		SSHCommand: wissh.SSHCommand{
			Command:    `sh -c 'for f in /mnt/boot/system-connections/*; do cat $f; echo; echo \#@@@@\#$f; done'`,
			IP:         ip,
			Port:       port,
			SSHKeyFile: sshKeyFile,
		},
	}
}

func (c *staticIPConfig) Run() error {
	if err := c.SSHCommand.Run(); err != nil {
		return err
	}

	// Don't show passwords in our output
	c.maskPasswords()

	configs, err := c.splitNetworkManagerConfigs()
	if err != nil {
		return err
	}
	c.passed = true

	for _, config := range configs {
		// This will set c.passed to false in case of errors.
		if err := c.checkStaticIPConfig(config); err != nil {
			return err
		}
	}

	return nil
}

func (c *staticIPConfig) Name() string {
	return "Verify Static IP Configuration"
}

func (c *staticIPConfig) Passed() bool {
	return c.passed
}

func (c *staticIPConfig) IlluminatingRemarks() (bool, string) {
	if !c.foundAnyStatic {
		return true, staticIPConfigRemarksNoStaticFound
	}
	if c.Passed() {
		return true, staticIPConfigRemarksSuccess + c.additionalRemarks
	}
	return true, staticIPConfigRemarksFailure + c.additionalRemarks
}

func (c *staticIPConfig) checkStaticIPConfig(config string) error {
	reFileName := regexp.MustCompile("(?m)^#@@@@#(.+)$")
	matches := reFileName.FindStringSubmatch(config)
	if len(matches) != 2 {
		return fmt.Errorf("unexpected matches when getting the config file name: %v", matches)
	}
	fileName := matches[0]
	cfg, err := ini.Load([]byte(config))
	if err != nil {
		return err
	}

	if err := c.checkStaticIPConfigForSection("ipv4", fileName, cfg); err != nil {
		return err
	}
	if err := c.checkStaticIPConfigForSection("ipv6", fileName, cfg); err != nil {
		return err
	}

	return nil
}

func (c *staticIPConfig) checkStaticIPConfigForSection(section, fileName string, cfg *ini.File) error {
	method := cfg.Section(section).Key("method").String()
	ip := cfg.Section(section).Key("address1").String()
	fileOnly := path.Base(fileName)
	if method == "manual" && ip == "" {
		c.foundAnyStatic = true
		c.passed = false
		c.additionalRemarks += fmt.Sprintf("* Problem on `%v` (section `[%v]`): `method` is set to `manual`, but the IP address is not set.", fileOnly, section)
	} else if ip != "" && method != "manual" {
		c.foundAnyStatic = true
		c.passed = false
		c.additionalRemarks += fmt.Sprintf("* Problem on `%v` (section `[%v]`): The IP address is set to `%v`, but the `method` is not set to `manual`, but the IP address is not set.", fileOnly, section, ip)
	} else if ip != "" && method == "manual" {
		c.foundAnyStatic = true
		c.additionalRemarks += fmt.Sprintf("* On `%v` (section `[%v]`): The IP address is correctly set to `%v`.", fileOnly, section, ip)
	}
	return nil
}

// splitNetworkManagerConfigs takes the output from the command ran over SSH and
// splits it into separate config files. Ignore files that shall be ignored.
func (c *staticIPConfig) splitNetworkManagerConfigs() ([]string, error) {
	result := []string{}
	reFile := regexp.MustCompile("^#@@@@#.+$")
	reIgnore := regexp.MustCompile(".ignore$")
	builder := strings.Builder{}
	scanner := bufio.NewScanner(strings.NewReader(c.StdOut))

	for scanner.Scan() {
		line := scanner.Text()
		if _, err := builder.WriteString(line); err != nil {
			return nil, err
		}
		if _, err := builder.WriteRune('\n'); err != nil {
			return nil, err
		}
		if reFile.MatchString(line) {
			if !reIgnore.MatchString(line) {
				result = append(result, builder.String())
			}
			builder.Reset()
		}
	}

	return result, nil
}

// maskPasswords masks WiFi network passwords from c.StdOut.
func (c *staticIPConfig) maskPasswords() {
	re := regexp.MustCompile(`(psk=)[^\n]*`)
	c.StdOut = re.ReplaceAllString(c.StdOut, "$1*****")
}

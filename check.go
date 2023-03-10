package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Check defines how any of the tests we can run looks like.
type Check interface {

	// Name returns a user-friendly name for the check. This is meant to be
	// displayed in the UI to identify the check and should give the user an
	// idea of what it does.
	Name() string

	// Run runs the test. The return value tells if there was some error while
	// running the Check, not if the check has passed.
	//
	// A Check can assume that Run will be the first of its methods to be
	// called. Furthermore, if the error is not nil, it can assume the other
	// methods will not be called.
	Run() error

	// Passed checks if the test passed.
	Passed() bool

	// IlluminatingRemarks returns any remarks that may illuminate the user.
	// Because here we are not satisfied with simple feedback: we explain why
	// things are failing, how balena works, what the user could try to do to
	// solve the issue, and, when appropriate, help with transcendence.
	//
	// The first return value tells if the check actually has remarks. The
	// second one contains the remarks themselves, as a Markdown string.
	IlluminatingRemarks() (bool, string)

	// Details returns details about the execution of the check. This should
	// include all the low-level details a more advanced user could make good
	// use of.
	//
	// The first return value tells if the check actually has details. The
	// second one contains the details themselves, as a Markdown string.
	Details() (bool, string)
}

// SSHCommand provides some boilerplate for implementing a Check that works by
// running a SSH command on the device.
type SSHCommand struct {
	// The IP address of the device in which we'll run the SSH command. Must be
	// set before calling Run.
	IP string

	// The port to which we'll connect to run the SSH command. Must be set
	// before calling Run.
	Port string

	// The path to the file with the private key to use when authenticating with
	// the device. Must be set before calling Run.
	//
	// TODO: This is required for now, and must point to a valid file. This
	// doesn't prevent us from connecting to devices in development mode (which
	// don't need authentication by default), but is kinda odd. Should be
	// optional.
	SSHKeyFile string

	// Command contains the command to run over SSH. Must be set before calling
	// Run.
	Command string

	// ExitStatus contains the exit status code of the SSH command. This is set
	// by Run.
	ExitStatus int

	// StdOut contains the standard output of the SSH command. This is set by
	// Run.
	StdOut string

	// StdErr contains the standard error of the SSH command. This is set by
	// Run.
	StdErr string
}

func (c *SSHCommand) Run() error {
	// Run the command
	runner, err := NewSSHRunner("root", c.IP+":"+c.Port, c.SSHKeyFile)
	if err != nil {
		return fmt.Errorf("preparing to run ssh command: %w", err)
	}
	defer runner.Destroy()

	c.StdOut, c.StdErr, err = runner.Run(c.Command)
	if err != nil {
		return fmt.Errorf("running ssh command: %w", err)
	}

	// Get the status code from the SSH command. We programmatically discard the
	// trailing line end instead of using `echo -n` just be avoid issues with
	// some hypothetical (maybe old) version of `echo` that can't handle `-n`.
	stdOut, _, err := runner.Run("echo $?")
	if err != nil {
		return fmt.Errorf("getting ssh command exit status: %w", err)
	}
	stdOut = strings.Split(stdOut, "\n")[0]

	c.ExitStatus, err = strconv.Atoi(stdOut)
	if err != nil {
		return fmt.Errorf("interpreting exit status (%q): %w", stdOut, err)
	}

	return nil
}

func (c *SSHCommand) Details() (bool, string) {
	return true, fmt.Sprintf("**SSH command:**\n\n`%v`\n\n", c.Command) +
		fmt.Sprintf("**Exit status:**\n\n`%v`\n\n", c.ExitStatus) +
		fmt.Sprintf("**Standard output:**\n\n```\n%v\n```\n\n", c.StdOut) +
		fmt.Sprintf("**Standard error:**\n\n```\n%v\n```\n\n", c.StdErr)
}

package wissh

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

// SSHRunner allows to run SSH commands. Technically, it's a wrapper around Go's
// ssh.Client.
type SSHRunner struct {
	client *ssh.Client
}

// NewSSHRunner creates a new SSHRunner. It will try to start an SSH session by
// user, on a server located at addr, and autheticating using the key at
// keyFile. Note that addr shall include the port, like "10.0.0.1:22222").
//
// You must call Destroy() on the returned SSHRunner when it is no longer
// needed.
func NewSSHRunner(user, addr, keyFile string) (*SSHRunner, error) {
	pk, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("reading private key file: %q", err)
	}
	signer, err := ssh.ParsePrivateKey(pk)
	if err != nil {
		return nil, fmt.Errorf("getting signer from key: %q", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // TODO: Add proper host key verification
	}

	config.Auth = append(config.Auth, nil)

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return nil, fmt.Errorf("creating ssh runner client: %q", err)
	}

	return &SSHRunner{
		client: client,
	}, nil
}

// Destroy frees all resources used by the Runner.
func (s *SSHRunner) Destroy() {
	s.client.Close()
}

// Run runs the cmd command over SSH.
func (s *SSHRunner) Run(cmd string) (stdOut string, stdErr string, err error) {
	session, err := s.client.NewSession()
	if err != nil {
		return "", "", fmt.Errorf("creating ssh runner session: %q", err)
	}
	defer session.Close()

	var stdoutBuff bytes.Buffer
	var stderrBuff bytes.Buffer
	session.Stdout = &stdoutBuff
	session.Stderr = &stderrBuff

	if err := session.Run(cmd); err != nil {
		return "", "", fmt.Errorf("running ssh command '%v': %q", cmd, err)
	}
	return stdoutBuff.String(), stderrBuff.String(), nil
}

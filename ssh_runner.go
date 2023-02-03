package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

type SSHRunner struct {
	client *ssh.Client
}

func NewSSHRunner(user, addr string) (*SSHRunner, error) {
	pk, err := os.ReadFile("/home/lmb/.ssh/id_rsa_wissh_test") // TODO: Make this configurable
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

func (s *SSHRunner) Destroy() {
	s.client.Close()
}

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

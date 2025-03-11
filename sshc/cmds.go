package sshc

import (
	"fmt"
	"os"
	"ssha/models"

	"golang.org/x/crypto/ssh"
)

func SSHintoHostViaPassword(host models.Host) error {
	config := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(host.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Insecure, but convenient for testing. Use with caution!
	}

	addr := fmt.Sprintf("%s:%d", host.Hostname, host.Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return fmt.Errorf("failed to dial SSH: %w", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return fmt.Errorf("request for pseudo terminal failed: %w", err)
	}

	if err := session.Shell(); err != nil {
		return fmt.Errorf("failed to start shell: %w", err)
	}

	err = session.Wait() // wait for the session to finish.
	if err != nil {
		return fmt.Errorf("session exited with error: %w", err)
	}

	return nil
}

func SSHintoHostViaPrivateKey(host models.Host) error {

	keyPath := host.PrivateKey

	key, err := os.ReadFile(keyPath)
	if err != nil {
		return fmt.Errorf("failed to read private key: %w", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %w", err)
	}

	config := &ssh.ClientConfig{
		User: host.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	addr := fmt.Sprintf("%s:%d", host.Hostname, host.Port)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		return fmt.Errorf("failed to dial SSH: %w", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return fmt.Errorf("request for pseudo terminal failed: %w", err)
	}

	if err := session.Shell(); err != nil {
		return fmt.Errorf("failed to start shell: %w", err)
	}

	err = session.Wait() // wait for the session to finish.
	if err != nil {
		return fmt.Errorf("session exited with error: %w", err)
	}

	return nil
}

package motdsource

import (
	"fmt"
	"io"
	"io/ioutil"

	motdv1alpha1 "github.com/alanraison/motd-operator/pkg/apis/motd/v1alpha1"

	"golang.org/x/crypto/ssh"
)

func fetchMotd(options motdv1alpha1.MotdSourceSpec) (string, error) {
	signer, err := ssh.ParsePrivateKey([]byte(options.PrivateKey))
	if err != nil {
		return "", fmt.Errorf("reading private key: %w", err)
	}

	c, err := ssh.Dial("tcp", options.Address, &ssh.ClientConfig{
		User: options.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return "", fmt.Errorf("connecting to %v: %w", options.Address, err)
	}
	sess, err := c.NewSession()
	if err != nil {
		return "", fmt.Errorf("starting session: %w", err)
	}
	defer sess.Close()

	out, err := sess.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("opening stdout: %w", err)
	}
	if err := sess.Run("cat /etc/motd"); err != nil {
		return "", fmt.Errorf("running command: %w", err)
	}
	lr := io.LimitReader(out, 1024)
	bs, err := ioutil.ReadAll(lr)
	if err != nil {
		return "", fmt.Errorf("reading command output: %w", err)
	}
	return string(bs), nil
}

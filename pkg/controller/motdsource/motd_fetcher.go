package motdsource

import (
	"fmt"
	"io"
	"io/ioutil"
	"time"

	motdv1alpha1 "github.com/alanraison/motd-operator/pkg/apis/motd/v1alpha1"
	"github.com/pkg/errors"

	"golang.org/x/crypto/ssh"
)

func fetchMotd(options motdv1alpha1.MotdSourceSpec) (time.Time, string, error) {
	signer, err := ssh.ParsePrivateKey([]byte(options.PrivateKey))
	if err != nil {
		return time.Time{}, "", errors.Errorf("reading private key: %w", err)
	}

	c, err := ssh.Dial("tcp", fmt.Sprint(options.Hostname, ":", options.Port), &ssh.ClientConfig{
		User: options.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	})
	if err != nil {
		return time.Time{}, "", errors.Errorf("connecting to %v: %w", options.Hostname, err)
	}
	sess, err := c.NewSession()
	if err != nil {
		return time.Time{}, "", errors.Errorf("starting session: %w", err)
	}
	defer sess.Close()

	out, err := sess.StdoutPipe()
	if err != nil {
		return time.Time{}, "", errors.Errorf("opening stdout: %w", err)
	}
	if err := sess.Run("cat /etc/motd"); err != nil {
		return time.Time{}, "", errors.Errorf("running command: %w", err)
	}
	lr := io.LimitReader(out, 1024)
	bs, err := ioutil.ReadAll(lr)
	if err != nil {
		return time.Time{}, "", errors.Errorf("reading command output: %w", err)
	}
	return time.Now(), fmt.Sprint(bs), nil
}

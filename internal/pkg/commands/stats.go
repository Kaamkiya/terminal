package commands

import (
	"fmt"
	
	"github.com/charmbracelet/ssh"
)

func statsCmd(session ssh.Session) {
	fmt.Fprintln(
		session,
		"User:",
		session.User(),
	)
	if session.PublicKey() != nil {
		fmt.Fprintln(
			session, 
			"SSH public key type:", 
			session.PublicKey().Type(),
		)
		fmt.Fprintln(
			session,
			"SSH public key:",
			string(session.PublicKey().Marshal()),
		)
	} else {
		fmt.Fprintln(
			session,
			"SSH public key:",
			"no public key used",
		)
	}
	fmt.Fprintln(
		session,
		"Environment:",
		session.Environ(),
	)
	fmt.Fprintln(
		session,
		"Connected from:",
		session.RemoteAddr().String(),
	)
}

package commands

import (
	"fmt"
	"time"
	
	"github.com/charmbracelet/ssh"
)

func statsCmd(session ssh.Session, joinTime time.Time) {
	// Print the name the user joined as.
	fmt.Fprintln(
		session,
		"User:",
		session.User(),
	)
	
	// Print the user's public key, or "no public key used" if they don't have one.
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
	
	// Print environment variables.
	fmt.Fprintln(
		session,
		"Environment:",
		session.Environ(),
	)

	// Print the IP the user connected from.
	fmt.Fprintln(
		session,
		"Connected from:",
		session.RemoteAddr().String(),
	)

	// Print the amount of time the user has been connected.
	fmt.Fprintln(
		session,
		"Time connected:",
		time.Since(joinTime).String(),
	)
}

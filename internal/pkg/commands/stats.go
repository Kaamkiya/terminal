package commands

import (
	"fmt"
	"time"
	
	"codeberg.org/Kaamkiya/terminal/internal/pkg/style"

	"github.com/charmbracelet/ssh"
)

func statsCmd(session ssh.Session, joinTime time.Time, styles style.Style) {
	// Print the name the user joined as.
	fmt.Fprintln(
		session,
		styles.Gray.Render("User:"),
		session.User(),
	)
	
	// Print the user's public key, or "no public key used" if they don't have one.
	if session.PublicKey() != nil {
		fmt.Fprintln(
			session, 
			styles.Gray.Render("SSH public key type:"), 
			session.PublicKey().Type(),
		)
		fmt.Fprintln(
			session,
			styles.Gray.Render("SSH public key:"),
			string(session.PublicKey().Marshal()),
		)
	} else {
		fmt.Fprintln(
			session,
			styles.Gray.Render("SSH public key:"),
			"no public key used",
		)
	}
	
	// Print environment variables.
	fmt.Fprintln(
		session,
		styles.Gray.Render("Environment:"),
		session.Environ(),
	)

	// Print the IP the user connected from.
	fmt.Fprintln(
		session,
		styles.Gray.Render("Connected from:"),
		session.RemoteAddr().String(),
	)

	// Print the amount of time the user has been connected.
	fmt.Fprintln(
		session,
		styles.Gray.Render("Time connected:"),
		time.Since(joinTime).String(),
	)
}

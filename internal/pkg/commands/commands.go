package commands

import (
	"fmt"
	"io"
	"time"

	"codeberg.org/Kaamkiya/terminal/internal/pkg/style"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"

	"golang.org/x/term"
)

func CommandLine(session ssh.Session, styles style.Style) {
	joinTime := time.Now()

	prompt := fmt.Sprintf(
		"%s@%s$ ",
		styles.Green.Render(session.User()),
		styles.Blue.Render(session.LocalAddr().String()),
	)

	terminal := term.NewTerminal(session, prompt)

	// Just so the user knows what they're doing when they connect.
	helpCmd(session, styles)

	for {
		input, err := terminal.ReadLine()
		if err == io.EOF {
			fmt.Fprintln(session, "Have a nice day :)")
			break
		}

		if err != nil {
			log.Error("Failed to read input", "error", err)
			fmt.Fprintln(session, "Failed to read input, exiting")
			break
		}

		switch input {
		case "about":
			aboutCmd(session)
		case "projects":
			projectsCmd(session, styles)
		case "exit":
			fmt.Fprintln(session, "Have a nice day :)")
			return
		case "stats":
			statsCmd(session, joinTime, styles)
		case "contact":
			contactCmd(session)
		case "help":
			helpCmd(session, styles)
		default:
			fmt.Fprintln(session, styles.Red.Render("Invalid command. Type help for a list of commands."))
		}
	}
}

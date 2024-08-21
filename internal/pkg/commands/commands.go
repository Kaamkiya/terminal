package commands

import (
	"fmt"
	"io"

	"codeberg.org/Kaamkiya/terminal/internal/pkg/style"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"

	"golang.org/x/term"
)

func CommandLine(session ssh.Session, style style.Style) {
	prompt := fmt.Sprintf(
		"%s@%s$ ",
		style.Green.Render(session.User()),
		style.Blue.Render(session.LocalAddr()),
	)
	terminal := term.NewTerminal(session, prompt)

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
			projectsCmd(session)
		case "exit":
			fmt.Fprintln(session, "Have a nice day :)")
			return
		case "stats":
			statsCmd(session)
		//case "uptime":
			//uptimeCmd(session)
		default:
			fmt.Fprintln(session, "Invalid command. Type help for a list of commands.")
		}
	}
}

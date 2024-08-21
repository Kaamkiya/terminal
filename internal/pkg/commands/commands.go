package commands

import (
	"fmt"
	"io"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"

	"golang.org/x/term"
)

func CommandLine(session ssh.Session) {
	prompt := fmt.Sprintf(
		"%s@%s$ ",
		session.User(),
		session.LocalAddr(),
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
		/*case "stats":
			statsCmd(session)
		case "uptime":
			uptimeCmd(session)
		*/
		default:
			fmt.Fprintln(session, "Invalid command. Type help for a list of commands.")
		}
	}
}

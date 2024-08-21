package commands

import (
	"fmt"
	"text/tabwriter"

	"codeberg.org/Kaamkiya/terminal/internal/pkg/style"

	"github.com/charmbracelet/ssh"
)

func helpCmd(session ssh.Session, styles style.Style) {
	basicCommands := map[string]string{
		"about": "some stuff about me",
		"projects": "my top 5 most recent codeberg repos",
		"stats": "show some statistics about you",
		"help": "show this help message",
	}

	dangerCommands := map[string]string{
		"exit": "leave this terminal session",
	}

	tw := tabwriter.NewWriter(session, 8, 16, 4, '\t', 0)

	for cmd, help := range basicCommands {
		fmt.Fprintln(tw, fmt.Sprintf(
			"%s\t%s",
			styles.Green.Render(cmd),
			help,
		))
	}

	for cmd, help := range dangerCommands {
		fmt.Fprintln(tw, fmt.Sprintf(
			"%s\t%s",
			styles.Red.Render(cmd),
			help,
		))
	}
	tw.Flush()
}

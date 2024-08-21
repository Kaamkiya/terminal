package commands

import (
	"fmt"

	"codeberg.org/Kaamkiya/terminal/internal/pkg/style"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/lipgloss/table"
)

func helpCmd(session ssh.Session, styles style.Style) {
	headers := []string{"Name", "Description"}
	
	commands := map[string]string{
		"about": "some stuff about me",
		"projects": "my top 5 most recent codeberg repos",
		"stats": "show some statistics about you",
		"help": "show this help message",
		"exit": "leave the terminal session",
	}
	rows := [][]string{}

	for cmd, help := range commands {
		rows = append(rows, []string{
			styles.Green.Render(cmd),
			help,
		})
	}

	t := table.New().
		Headers(headers...).
		Rows(rows...).
		Render()
	
	fmt.Fprintln(session, t)
}

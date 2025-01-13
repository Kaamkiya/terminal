package commands

import (
	"fmt"

	"codeberg.org/Kaamkiya/terminal/internal/style"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/ssh"
)

func projectsCmd(session ssh.Session, styles style.Style) {
	headers := []string{"Name", "Description", "Language", "URL"}
	data := [][]string{
		{"gg", "A collection of TUI games", "Go", "https://github.com/Kaamkiya/gg"},
		{"terminal", "This very program!", "Go", "https://github.com/Kaamkiya/terminal"},
	}

	t := table.New().
		Headers(headers...).
		Rows(data...).
		Render()

	fmt.Fprintln(session, t)
}

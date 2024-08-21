package style

import (
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/lipgloss"

	"github.com/muesli/termenv"
)

type Style struct {
	Renderer *lipgloss.Renderer
	Red lipgloss.Style
	Green lipgloss.Style
	Blue lipgloss.Style
}

func GetStyles(session ssh.Session) Style {
	rend := lipgloss.NewRenderer(session)
	rend.SetOutput(termenv.NewOutput(session))

	return style{
		Renderer: rend,
		Red: rend.NewStyle().Foreground(lipgloss.NewColor("#f32727")),
		Green: rend.NewStyle().Foreground(lipgloss.NewColor("#0a5c36"),
		Blue: rend.NewStyle().Foreground(lipgloss.NewColor("#0e49b5"),
	}
}

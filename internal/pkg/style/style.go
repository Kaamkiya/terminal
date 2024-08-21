package style

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"

	"github.com/muesli/termenv"
)

type Style struct {
	Renderer *lipgloss.Renderer
	Red      lipgloss.Style
	Green    lipgloss.Style
	Blue     lipgloss.Style
	Gray     lipgloss.Style
}

func GetStyles(session ssh.Session) Style {
	rend := lipgloss.NewRenderer(session)
	rend.SetOutput(termenv.NewOutput(session, termenv.WithUnsafe()))

	return Style{
		Renderer: rend,
		Red:      rend.NewStyle().Foreground(lipgloss.Color("#ed2828")),
		Green:    rend.NewStyle().Foreground(lipgloss.Color("#14e87e")),
		Blue:     rend.NewStyle().Foreground(lipgloss.Color("#147ee8")),
		Gray:     rend.NewStyle().Foreground(lipgloss.Color("#aaaaaa")),
	}
}

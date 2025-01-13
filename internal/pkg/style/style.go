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
	r := lipgloss.NewRenderer(session)
	r.SetOutput(termenv.NewOutput(session, termenv.WithUnsafe()))

	return Style{
		Renderer: r,
		Red:      r.NewStyle().Foreground(lipgloss.Color("#ed2828")),
		Green:    r.NewStyle().Foreground(lipgloss.Color("#14e87e")),
		Blue:     r.NewStyle().Foreground(lipgloss.Color("#147ee8")),
		Gray:     r.NewStyle().Foreground(lipgloss.Color("#aaaaaa")),
	}
}

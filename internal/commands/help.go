package commands

import (
	"fmt"

	"github.com/charmbracelet/ssh"
)

func helpCmd(s ssh.Session) {
	msg := `
about       some stuff about me"
projects    projects I'm proud of"
stats       some statistics about you"
contact     how to contact me"
help        show this help message"
exit        leave the terminal"
`

	fmt.Fprintln(s, msg)
}

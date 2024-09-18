package commands

import (
	"time"
	
	"codeberg.org/Kaamkiya/terminal/internal/pkg/animate"

	"github.com/charmbracelet/ssh"
)

func contactCmd(session ssh.Session) {
	animate.TypeWriter(
		session,
		8*time.Millisecond,
		"You can reach me on GitHub at https://github.com/Kaamkiya\n",
	)
}

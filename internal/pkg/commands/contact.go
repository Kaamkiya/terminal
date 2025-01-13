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
		`You can reach me on...

* Github:    Kaamkiya
* Mastodon:  @nonexistent@hachyderm.io
* Lemmy:     @kaamkiy@lemmy.ml
* Hachyderm: kaamkiya
* Daily.dev: kaamkiya
* Dev.to:    kaamkiya

My ideal method of contact is Mastdon.
`,
	)
}

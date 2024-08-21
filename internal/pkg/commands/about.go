package commands

import (
	"fmt"
	"time"

	"codeberg.org/Kaamkiya/terminal/internal/pkg/animate"

	"github.com/charmbracelet/ssh"
)

func aboutCmd(session ssh.Session) {
	year := time.Now().Year()

	animate.TypeWriter(
		session,
		time.Millisecond,
		fmt.Sprintf(`
Hello! I'm Kaamkiya.

I am a human who happens to like coding, reading, and other miscellaneous
things.

I like to use my coding skills to make fun programs (like this one) that can,
sometimes, end up helping people. Almost everything I code is FLOSS software,
so you can copy it, contribute to it, and more.

I like contributing to FLOSS software too, especially when it's in Go.

By the way, this program is FLOSS, so you can make something like it if you
want to!

Source code: https://codeberg.org/Kaamkiya/terminal

More stuff about me:
- Pronouns: he/him
- Age: somewhere between 10 and 80
- OS: Linux Mint 22
- Human languages: English (fluent), French (fairly fluent), and learning more.
- Skills: backend and basic web dev in Go (%dy), full stack Python (%dy), full
          stack TypeScript using Deno and Fresh (%dy).
- Favorite coding language: Go. I don't know why, I just really love it.`,
		year-2023,
		year-2022,
		year-2023,
		),
	)
}

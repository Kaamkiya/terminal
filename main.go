package main

import (
	"errors"
	"time"
	"net"
	"flag"
	"strconv"

	"codeberg.org/Kaamkiya/terminal/internal/pkg/commands"
	"codeberg.org/Kaamkiya/terminal/internal/pkg/style"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
)

var (
	flagHost = flag.String("host", "0.0.0.0", "where to host the server")
	flagPort = flag.Int("port", 2222, "the port to use")
	flagIdleTimeout = flag.Duration("idletimeout", 15*time.Minute, "idle timeout for connections")
)

const banner = `  _        _ _     
 | |_  ___| | |___ 
 | ' \/ -_) | / _ \
 |_||_\___|_|_\___/
`

func main() {
	hostURL := net.JoinHostPort(*flagHost, strconv.Itoa(*flagPort))

	server, err := wish.NewServer(
		wish.WithAddress(hostURL),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(session ssh.Session) {
					styles := style.GetStyles(session)
					commands.CommandLine(session, styles)
				}
			},
			logging.Middleware(),
		),
		wish.WithBanner(banner),
		wish.WithIdleTimeout(*flagIdleTimeout),
	)
	if err != nil {
		log.Error("Failed to start server", "error", err)
	}

	log.Info("Server started on " + hostURL)
	if err = server.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Failed to start server", "error", err)
	}
}


package app

import (
	"errors"
	"net"
	"strconv"

	"codeberg.org/Kaamkiya/terminal/internal/pkg/commands"
	"codeberg.org/Kaamkiya/terminal/internal/pkg/conf"
	"codeberg.org/Kaamkiya/terminal/internal/pkg/style"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
)

func Run() {
	config := conf.LoadConfig("config.yaml")
	hostURL := net.JoinHostPort(config.Host, strconv.Itoa(config.Port))

	server, err := wish.NewServer(
		wish.WithAddress(hostURL),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(session ssh.Session) {
					styles := style.GetStyles(session)
					wish.Println(session, styles.Green.Render("Welcome to my terminal!"))
					commands.CommandLine(session, styles)
				}
			},
			logging.Middleware(),
		),
		wish.WithBanner(getBanner(config.BannerPath)),
		wish.WithIdleTimeout(config.IdleTimeout),
	)
	if err != nil {
		log.Error("Failed to start server", "error", err)
	}

	log.Info("Server started on " + hostURL)
	if err = server.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Failed to start server", "error", err)
	}
}

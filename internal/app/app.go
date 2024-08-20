package app

import (
	"net"

	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
)

const (
	host = "localhost"
	port = "18187"
)

func Run() {
	server, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(session ssh.Session) {
					wish.Println(session, "Welcome to my terminal!")
				}
			},
			logging.Middleware(),
		),
		wish.WithBanner(getBanner()),
	)
	if err != nil {
		log.Error("Failed to start server", "error", err)
	}

	if err = server.ListenAndServe(); err != nil {
		if err == ssh.ErrServerClosed {
			log.Info("Server closed")
		} else {
			log.Error("Failed to start server", "error", err)
		}
	}
}

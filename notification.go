package main

import (
	"fmt"
	"github.com/primaldarkness/mud-redirector-service/listener"
	log "github.com/sirupsen/logrus"
	"net"
)

type BannerCommand struct {
	RedirectTo string `arg:"true" help:"Redirect connections to a specific address and port (ie mud.example.com:5000)."`
}

func (cmd *BannerCommand) Run(cli *CliContext) error {
	ctx := cli.ctx
	logger := ctx.Value(ctxLogger).(*log.Entry)

	l := listener.Listener{Addr: cli.Listen}

	logger.Infof("will redirect connections to %s", cli.Redirect.RedirectTo)

	return l.Run(ctx, func(conn net.Conn) {
		logger = logger.WithField("remoteAddr", conn.RemoteAddr())
		_, _ = fmt.Fprint(conn, notificationBanner(cli))
		logger.Debugf("remote connection %s has been established", conn.RemoteAddr())

		conn.Close()
	})
}

func notificationBanner(cli *CliContext) string {
	return fmt.Sprintf(`Welcome to %s. Our MUD has moved to %s. 
Please update your connection and try again.
`, cli.MudName, cli.Banner.RedirectTo)

}

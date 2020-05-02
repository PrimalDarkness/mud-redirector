package main

import (
	"context"
	"fmt"
	"github.com/primaldarkness/mud-redirector-service/listener"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"time"
)

type RedirectCommand struct {
	RedirectTo string `arg:"true" help:"Redirect connections to a specific address and port (ie mud.example.com:5000)."`
}

func (cmd *RedirectCommand) Run(cli *CliContext) error {
	ctx := cli.ctx
	logger := ctx.Value(ctxLogger).(*log.Entry)

	l := listener.Listener{Addr: cli.Listen}

	logger.Infof("will redirect connections to %s", cli.Redirect.RedirectTo)

	return l.Run(ctx, func(conn net.Conn) {
		ctx, cancel := context.WithCancel(ctx)

		logger = logger.WithField("remoteAddr", conn.RemoteAddr())
		_, _ = fmt.Fprint(conn, redirectBanner(cli))
		logger.Debugf("remote connection %s has been established", conn.RemoteAddr())

		time.Sleep(3 * time.Second)

		_, _ = fmt.Fprintf(conn, "Proxying to remote %s...\n", cli.Redirect.RedirectTo)

		rConn, err := net.Dial("tcp", cli.Redirect.RedirectTo)
		if err != nil {
			logger.Errorf("could not connect to %s", cli.Redirect.RedirectTo)
			_, _ = fmt.Fprintf(conn, "could not connect to %s", cli.Redirect.RedirectTo)
		}

		defer conn.Close()
		defer rConn.Close()

		logger.Infof("proxy has been established")
		go func() {
			io.Copy(conn, rConn)
			cancel()
		}()
		go func() {
			io.Copy(rConn, conn)
			cancel()
		}()

		<-ctx.Done()

		logger.Infof("connection has closed")
	})
}

func redirectBanner(cli *CliContext) string {
	return fmt.Sprintf(`
Welcome to %s. Our MUD has moved to %s, but we'll go ahead
and connect you now.
`, cli.MudName, cli.Redirect.RedirectTo)
}

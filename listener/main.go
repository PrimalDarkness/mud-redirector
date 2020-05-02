package listener

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net"
)

type Listener struct {
	Addr string
}

func (l *Listener) Run(ctx context.Context, handler func(conn net.Conn)) error {
	logger := ctx.Value("logger").(*log.Entry)
	ln, err := net.Listen("tcp", l.Addr)
	if err != nil {
		return err
	}

	logger.Debugf("service is listening on addr %s", l.Addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handler(conn)
	}
}

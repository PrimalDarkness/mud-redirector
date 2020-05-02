package main

import (
	"context"
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

type CliContext struct {
	*CLI
	ctx context.Context
}

type CLI struct {
	Verbose  int             `type:"counter" short:"v" help:"Tweak the verbosity of the logs."`
	Listen   string          `short:"l" required:"true" help:"The address for the service to listen on (0.0.0.0:5000)"`
	MudName  string          `short:"n" required:"true" help:"The name of your MUD."`
	Redirect RedirectCommand `cmd:"true" help:"Proxy connecting user to the MUD (warning, IP information will be hidden)."`
	Banner   BannerCommand   `cmd:"true" help:"Display a banner to the connecting user and disconnect immediately."`
}

const (
	ctxLogger = "logger"
	ctxCLI    = "cli"
)

func main() {
	var cli CLI
	kCtx := kong.Parse(&cli,
		kong.Name("mud-redirector"),
		kong.Description("A simple notification and/or proxy service for redirecting MUD connections."),
		kong.UsageOnError(),
		kong.Configuration(kong.JSON, "~/.config.json"),
	)

	logger := setupLogger(cli.Verbose)

	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxLogger, logger)
	ctx = context.WithValue(ctx, ctxCLI, cli)

	err := kCtx.Run(&CliContext{ctx: ctx, CLI: &cli})
	if err != nil {
		logger.Fatal(err)
	}
}

func setupLogger(verbosity int) *log.Entry {
	logger := log.New()

	switch verbosity {
	case 0:
		logger.SetLevel(log.InfoLevel)
	case 1:
		logger.SetLevel(log.DebugLevel)
	case 2:
		fallthrough
	default:
		logger.SetLevel(log.TraceLevel)
	}

	return log.NewEntry(logger)
}

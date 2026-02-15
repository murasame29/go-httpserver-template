package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"strings"

	"github.com/murasame29/go-httpserver-template/internal/config"
	"github.com/murasame29/go-httpserver-template/internal/container"
	"github.com/murasame29/go-httpserver-template/internal/server"
)

type envFlag []string

func (e *envFlag) String() string {
	return strings.Join(*e, ",")
}

func (e *envFlag) Set(v string) error {
	*e = append(*e, v)
	return nil
}

func parseLogLevel(l string) slog.Level {
	switch strings.ToLower(l) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func main() {
	var (
		envFile  envFlag
		logLevel string
	)

	flag.Var(&envFile, "e", "path to .env file (e.g. -e .env -e another.env)")
	flag.StringVar(&logLevel, "ll", "debug", "log level (debug, info, warn, error)")
	flag.Parse()

	slog.SetDefault(
		slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: parseLogLevel(logLevel),
		})),
	)

	if err := run(envFile); err != nil {
		slog.Error("failed to run application", "error", err)
		os.Exit(1)
	}
}

func run(envFiles []string) error {
	cfg, err := config.Load(envFiles...)
	if err != nil {
		return err
	}

	if err := container.New(cfg); err != nil {
		return err
	}

	srv, err := container.Invoke[*server.Server]()
	if err != nil {
		return err
	}

	return srv.RunWithGracefulShutdown(context.Background())
}

package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"github.com/murasame29/go-httpserver-template/cmd/config"
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
	switch {
	case strings.EqualFold(l, "debug") || strings.EqualFold(l, "DEBUG"):
		return slog.LevelDebug
	case strings.EqualFold(l, "info") || strings.EqualFold(l, "INFO"):
		return slog.LevelInfo
	case strings.EqualFold(l, "warn") || strings.EqualFold(l, "WARN"):
		return slog.LevelWarn
	case strings.EqualFold(l, "error") || strings.EqualFold(l, "ERROR"):
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func init() {
	// Usage: eg. go run main.go -e .env -e hoge.env -e fuga.env ...
	var (
		envFile  envFlag
		logLevel string
	)

	flag.Var(&envFile, "e", "path to .env file \n eg. -e .env -e another.env . ")
	flag.StringVar(&logLevel, "ll", "debug", "Change the loglevel. The default is debug .")
	flag.Parse()

	slog.SetDefault(
		slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: parseLogLevel(logLevel),
		})),
	)

	if err := config.LoadEnv(envFile...); err != nil {
		slog.Error("failed to laod env", "error", err)
	}
}

func main() {
	if err := run(); err != nil {
		slog.Error("failed to run application", "error", err)
	}
}

func run() error {
	// サーバーの起動
	if err := container.NewContainer(); err != nil {
		return err
	}

	handler, err := container.Invoke[http.Handler]()
	if err != nil {
		return err
	}

	server.New(
		handler,
		server.WithHost(config.Config.Server.Host),
		server.WithPort(config.Config.Server.Port),
	).RunWithGracefulShutdown(context.Background())

	return nil
}

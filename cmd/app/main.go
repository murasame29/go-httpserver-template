package main

import (
	"context"
	"flag"
	"net/http"
	"strings"

	"github.com/murasame29/go-httpserver-template/cmd/config"
	"github.com/murasame29/go-httpserver-template/internal/container"
	"github.com/murasame29/go-httpserver-template/internal/server"
	"github.com/murasame29/go-httpserver-template/pkg/log"
)

type envFlag []string

func (e *envFlag) String() string {
	return strings.Join(*e, ",")
}

func (e *envFlag) Set(v string) error {
	*e = append(*e, v)
	return nil
}

func init() {
	// Usage: eg. go run main.go -e .env -e hoge.env -e fuga.env ...
	var envFile envFlag
	flag.Var(&envFile, "e", "path to .env file \n eg. -e .env -e another.env . ")
	flag.Parse()

	if err := config.LoadEnv(envFile...); err != nil {
		log.Fatal(context.Background(), err, "message", "Error loading .env file")
	}
}
func main() {
	if err := run(); err != nil {
		log.Fatal(context.Background(), err)
	}
}

func run() error {
	// サーバーの起動
	if err := container.NewContainer(); err != nil {
		return err
	}

	// handler をinvoke
	var handler http.Handler
	if err := container.Invoke(func(h http.Handler) {
		handler = h
	}); err != nil {
		return err
	}

	server.New(
		handler,
		server.WithHost(config.Config.Server.Host),
		server.WithPort(config.Config.Server.Port),
		server.WithShutdownTimeout(config.Config.Server.ShutdownTimeout),
	).RunWithGracefulShutdown(context.Background())

	return nil
}

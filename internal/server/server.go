package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
)

const DefaultShutdownTimeout time.Duration = 10

type server struct {
	port            int
	host            string
	shutdownTimeout time.Duration

	l *logrus.Logger

	srv *http.Server
}

// Server はサーバーを表すインターフェースです。
type Server interface {
	Run() error
	Shutdown(ctx context.Context) error

	RunWithGracefulShutdown()
}

// New はサーバーを生成します。
func New(handler http.Handler, opts ...Option) Server {
	server := &server{
		port:            8080,
		host:            "localhost",
		shutdownTimeout: DefaultShutdownTimeout,
		l:               logrus.New(),
		srv:             new(http.Server),
	}

	for _, opt := range opts {
		opt(server)
	}

	server.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", server.host, server.port),
		Handler: handler,
	}

	return server
}

// Run はサーバーを起動します。
func (s *server) Run() error {
	s.l.Infof("server starting at %s", s.srv.Addr)
	return s.srv.ListenAndServe()
}

// Shutdown はサーバーを停止します。
func (s *server) Shutdown(ctx context.Context) error {
	s.l.Infof("server shutdown ...")
	return s.srv.Shutdown(ctx)
}

// RunWithGracefulShutdown はgraceful shutdownを行うサーバーを起動します。
func (s *server) RunWithGracefulShutdown() {
	go func() {
		if err := s.Run(); err != nil && err != http.ErrServerClosed {
			s.l.Fatalf("Listen And Serve error : %s", err.Error())
		}
	}()

	s.l.Infof("server starting at %s", s.srv.Addr)

	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit
	if err := s.Shutdown(ctx); err != nil {
		s.l.Fatalf("server shutdown error : %s", err.Error())
	}

	s.l.Infof("server shutdown completed")
}

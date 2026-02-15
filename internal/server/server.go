package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/murasame29/go-httpserver-template/internal/config"
	"golang.org/x/sync/errgroup"
)

// Server はHTTPサーバーを表します。
// pkg/lifecycle.Application インターフェースを実装しています。
type Server struct {
	cfg *config.ServerConfig
	srv *http.Server
}

// New はサーバーを生成します。
func New(cfg *config.ServerConfig, handler http.Handler) *Server {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return &Server{
		cfg: cfg,
		srv: srv,
	}
}

// Run はサーバーを起動します。
func (s *Server) Run(ctx context.Context) error {
	slog.InfoContext(ctx, "server starting", "addr", s.srv.Addr)
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve: %w", err)
	}
	return nil
}

// Shutdown はサーバーを停止します。
func (s *Server) Shutdown(ctx context.Context) error {
	slog.InfoContext(ctx, "server shutting down...")
	return s.srv.Shutdown(ctx)
}

// RunWithGracefulShutdown は graceful shutdown を行うサーバーを起動します。
func (s *Server) RunWithGracefulShutdown(ctx context.Context) error {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.Run(egCtx)
	})

	eg.Go(func() error {
		<-egCtx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
		defer cancel()

		return s.Shutdown(shutdownCtx)
	})

	if err := eg.Wait(); err != nil && err != context.Canceled {
		return err
	}

	slog.Info("server shutdown completed")
	return nil
}

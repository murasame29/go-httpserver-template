package server

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Option func(*Server)

// WithPort はポート番号を設定するオプションです。
func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

// WithHost はホスト名を設定するオプションです。
func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

// WithLogger はロガーを設定するオプションです。
func WithLogger(l *logrus.Logger) Option {
	return func(s *Server) {
		s.l = l
	}
}

// WithReadTimeout はリクエストの読み込みタイムアウトを設定するオプションです。
func WithReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.srv.ReadTimeout = timeout
	}
}

// WithWriteTimeout はレスポンスの書き込みタイムアウトを設定するオプションです。
func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.srv.WriteTimeout = timeout
	}
}

// WithIdleTimeout はアイドルタイムアウトを設定するオプションです。
func WithIdleTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.srv.IdleTimeout = timeout
	}
}

// WithShutdownTimeout はシャットダウンタイムアウトを設定するオプションです。
func WithShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}

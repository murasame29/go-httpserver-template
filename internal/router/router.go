package router

import (
	"net/http"

	"github.com/murasame29/go-httpserver-template/internal/middleware"
)

// NewServeMux は Go 1.22+ の net/http.ServeMux を利用した http.Handler を返します。
func NewServeMux() http.Handler {
	mux := http.NewServeMux()

	// Health Check endpoints
	mux.HandleFunc("GET /livez", handleLivez)
	mux.HandleFunc("GET /readyz", handleReadyz)

	// Middleware を適用
	handler := middleware.Chain(
		middleware.Recovery,
		middleware.Logging,
	)(mux)

	return handler
}

func handleLivez(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func handleReadyz(w http.ResponseWriter, r *http.Request) {
	// TODO: DB接続などの準備状態をチェック
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

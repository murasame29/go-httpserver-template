package middleware

import "net/http"

// Middleware は http.Handler をラップする関数型です。
type Middleware func(http.Handler) http.Handler

// Chain は複数のミドルウェアを連結します。
// 順序: 最初に渡したものが最も外側になります。
// 例: Chain(recovery, logging, tracing) → recovery(logging(tracing(handler)))
func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}

package lifecycle

import "context"

// Application はアプリケーションのライフサイクルを表すインターフェースです。
type Application interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

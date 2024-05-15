package log

import (
	"context"
	"io"
	"log/slog"
	"os"
)

// Logger は、slogをwrapしてる独自定義のLogger
type Logger struct {
	*slog.Logger
}

// New はslogをwrapしているLoggerを初期化する
func New(h slog.Handler) *Logger {
	l := &Logger{
		slog.New(h),
	}

	return l
}

// DefaultHandler は、slogのjsonハンドラーを返す
// 初期LogLevelはDebug
func DefaultHandler(w io.Writer, opts ...OptionFunc) slog.Handler {
	o := Option{
		Level: slog.LevelDebug,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return slog.NewJSONHandler(
		w,
		&slog.HandlerOptions{
			Level: o.Level,
		},
	)
}

// LoggerKey は、LoggerのKey
type LoggerKey struct{}

// FromContext は、contextからLoggerを取り出す
// 存在しない場合は、DefaultHandler(os.Stdout)で初期化される
func FromContext(ctx context.Context) *Logger {
	if l, ok := ctx.Value(LoggerKey{}).(*Logger); ok {
		return l
	}

	return New(DefaultHandler(os.Stdout))
}

// IntoContext は、contextにLoggerを格納する
func IntoContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, LoggerKey{}, logger)
}

// Debug は、LogレベルDebugを出力する
func Debug(ctx context.Context, msg string, keysAndValues ...any) {
	FromContext(ctx).Debug(msg, keysAndValues...)
}

// Info は、LogレベルInfoを出力する
func Info(ctx context.Context, msg string, keysAndValues ...any) {
	FromContext(ctx).Info(msg, keysAndValues...)
}

// Warn は、LogレベルWarnを出力する
func Warn(ctx context.Context, msg string, keysAndValues ...any) {
	FromContext(ctx).Warn(msg, keysAndValues...)
}

// Error は、LogレベルErrorを出力する
func Error(ctx context.Context, err error, keysAndValues ...any) {
	FromContext(ctx).Error(err.Error(), keysAndValues...)
}

// Fatal は、LogレベルErrorを出力し異常終了させる
func Fatal(ctx context.Context, err error, keysAndValues ...any) {
	FromContext(ctx).Error(err.Error(), keysAndValues...)
	os.Exit(1)
}

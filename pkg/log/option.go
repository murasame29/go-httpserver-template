package log

import "log/slog"

// Option はDefaultHandler用のOption
type Option struct {
	Level slog.Level
}

// OptionFunc は(ry
type OptionFunc func(*Option)

// WithLevel は、ログレベルを任意の値に指定する
func WithLevel(level slog.Level) OptionFunc {
	return func(o *Option) {
		o.Level = level
	}
}

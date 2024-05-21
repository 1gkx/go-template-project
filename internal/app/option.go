package app

import "log/slog"

type Option func(a *App)

func WithLogger(log *slog.Logger) Option {
	return func(a *App) {
		a.log = log
	}
}

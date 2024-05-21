package http

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
)

type Option func(srv *http.Server)

func WithReadTimeout(t time.Duration) Option {
	return func(srv *http.Server) {
		srv.ReadTimeout = t
	}
}

func WithWriteTimeout(t time.Duration) Option {
	return func(srv *http.Server) {
		srv.WriteTimeout = t
	}
}

func WithContext(ctx context.Context) Option {
	return func(srv *http.Server) {
		srv.BaseContext = func(_ net.Listener) context.Context {
			return ctx
		}
	}
}

func WithErrorLog(log *log.Logger) Option {
	return func(srv *http.Server) {
		srv.ErrorLog = log
	}
}

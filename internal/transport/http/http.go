package http

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
)

var errEmptyPort = errors.New("server port is required")

type Server struct {
	hs *http.Server
	l  net.Listener
}

func New(addr string, handler http.Handler, opts ...Option) (*Server, error) {
	var err error

	_, _, err = net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	for _, o := range opts {
		o(srv)
	}

	return &Server{
		hs: srv,
		l:  ln,
	}, nil
}

func (s *Server) ListerAndServe() error {
	return s.hs.Serve(s.l)
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("shutdown http server")
	return s.hs.Shutdown(ctx)
}

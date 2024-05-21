package grpc

import (
	"context"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type ServerRegistrator interface {
	RegisterService(sd *grpc.ServiceDesc, ss any)
}

type Registrator func(gs ServerRegistrator)

type Server struct {
	gs *grpc.Server
	l  net.Listener
}

func New(addr string, reg Registrator) (*Server, error) {
	var err error

	if _, _, err = net.SplitHostPort(addr); err != nil {
		return nil, err
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	srv := &Server{
		gs: grpc.NewServer(),
		l:  ln,
	}

	reg(srv.gs)

	return srv, nil
}

func (s *Server) ListerAndServe() error {
	return s.gs.Serve(s.l)
}

func (s *Server) Shutdown(ctx context.Context) error {
	slog.Info("shutdown grpc server")
	s.gs.GracefulStop()
	return nil
}

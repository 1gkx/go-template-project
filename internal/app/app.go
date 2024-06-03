package app

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"{{ ModuleName .UserName .ProjectName }}/internal/config"
	"{{ ModuleName .UserName .ProjectName }}/internal/transport/grpc"
	"{{ ModuleName .UserName .ProjectName }}/internal/transport/http"
	pbHealth "{{ ModuleName .UserName .ProjectName }}/pkg/api"
)

type App struct {
	log       *slog.Logger
	interrupt chan os.Signal
	shutdowns []func(context.Context) error
}

func New(opts ...Option) *App {
	app := &App{
		interrupt: make(chan os.Signal, 1),
	}

	for _, o := range opts {
		o(app)
	}

	return app
}

func (a *App) RegisterShutdown(fn func(ctx context.Context) error) {
	a.shutdowns = append(a.shutdowns, fn)
}

func (a *App) GracefulShutdown(ctx context.Context) error {
	gsCtx, gsCancel := context.WithTimeout(ctx, 5*time.Second)
	defer gsCancel()

	var err error
	for i := len(a.shutdowns) - 1; i >= 0; i-- {
		if err = a.shutdowns[i](gsCtx); err != nil {
			return err
		}
	}

	return ctx.Err()
}

func (a *App) Run() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var (
		err     error
		conf    *config.Config
		httpSrv *http.Server
		grpcSrv *grpc.Server
	)

	conf, err = config.Load()
	if err != nil {
		return err
	}

	httpSrv, err = http.New(conf.HTTP.Addr(), nil)
	if err != nil {
		return err
	}
	a.RegisterShutdown(httpSrv.Shutdown)

	grpcSrv, err = grpc.New(conf.GRPC.Addr(), func(gs grpc.ServerRegistrator) {
		pbHealth.RegisterHealthServiceServer(gs, new(pbHealth.UnimplementedHealthServiceServer))
	})
	if err != nil {
		return err
	}
	a.RegisterShutdown(grpcSrv.Shutdown)

	go signal.Notify(a.interrupt, syscall.SIGINT, syscall.SIGTERM)

	a.log.Info("Start application")

	select {
	case sig := <-a.interrupt:
		a.log.Debug("Received signal: %s", sig)
		return a.GracefulShutdown(ctx)
	}
}

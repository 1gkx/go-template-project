package config

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/ardanlabs/conf/v3"
)

const appEnvPrefix = "TMPL"

type Server struct {
	Host string `conf:"default:0.0.0.0"`
	Port int    `conf:"default:8080"`
}

func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

type Config struct {
	conf.Version
	HTTP Server
	GRPC Server
}

func Load() (*Config, error) {
	var cfg Config
	if h, err := conf.Parse(appEnvPrefix, &cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(h)
		}
		return nil, err
	}

	c, err := conf.String(&cfg)
	if err != nil {
		return nil, err
	}

	slog.Info("Config", slog.String("data", c))

	return &cfg, nil
}

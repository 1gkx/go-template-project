package config

import (
	"errors"
	"fmt"

	"github.com/ardanlabs/conf/v3"
)

const appEnvPrefix = ""

type Version struct {
	Build string
	Desc  string
}

type Config struct {
	Version Version
}

func Load() (*Config, error) {
	var cfg *Config
	if h, err := conf.Parse(appEnvPrefix, cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(h)
		}
		return nil, err
	}

	return cfg, nil
}

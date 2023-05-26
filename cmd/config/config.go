package config

import (
	"github.com/tkanos/gonfig"
	"log"
)

type Config struct {
	ApiPort     int    `env:"API_PORT"`
	PostgresURL string `env:"POSTGRES_URL"`
	DebugDb     bool   `env:"DEBUG_DB"`
}

func Load(path string) *Config {
	config := new(Config)

	if err := gonfig.GetConf(path, config); err != nil {
		log.Fatal("error to load configs: ", err)
	}

	return config
}

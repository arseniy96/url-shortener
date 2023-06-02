package config

import (
	"flag"
	"github.com/caarlos0/env"
)

type Options struct {
	Host         string `env:"SERVER_ADDRESS"`
	ResolveHost  string `env:"BASE_URL"`
	LoggingLevel string `env:"LOG_LEVEL"`
	Filename     string `env:"FILE_STORAGE_PATH"`
}

func InitConfig() *Options {
	options := &Options{}

	flag.StringVar(&options.Host, "a", "localhost:8080", "server host with port")
	flag.StringVar(&options.ResolveHost, "b", "http://localhost:8080", "resolve link address")
	flag.StringVar(&options.LoggingLevel, "l", "info", "log level")
	flag.StringVar(&options.Filename, "f", "/tmp/short-url-db.json", "storage file")
	flag.Parse()

	err := env.Parse(options)
	if err != nil {
		panic(err)
	}

	return options
}

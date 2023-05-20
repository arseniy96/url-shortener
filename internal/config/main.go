package config

import (
	"flag"
	"os"
)

var host *string
var resolveHost *string

type Options struct {
	Host        string
	ResolveHost string
}

func init() {
	host = flag.String("a", "localhost:8080", "server host with port")
	resolveHost = flag.String("b", "http://localhost:8080", "resolve link address")
}

func SetConfig() *Options {
	flag.Parse()

	if envServerAddr := os.Getenv("SERVER_ADDRESS"); envServerAddr != "" {
		host = &envServerAddr
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		resolveHost = &envBaseURL
	}

	return &Options{
		Host:        *host,
		ResolveHost: *resolveHost,
	}
}

package config

import (
	"flag"
	"os"
)

var host *string
var resolveHost *string
var loggingLevel *string
var filename *string

type Options struct {
	Host         string
	ResolveHost  string
	LoggingLevel string
	Filename     string
}

func init() {
	host = flag.String("a", "localhost:8080", "server host with port")
	resolveHost = flag.String("b", "http://localhost:8080", "resolve link address")
	loggingLevel = flag.String("l", "info", "log level")
	filename = flag.String("f", "/tmp/short-url-db.json", "storage file")
}

func SetConfig() *Options {
	flag.Parse()

	if envServerAddr := os.Getenv("SERVER_ADDRESS"); envServerAddr != "" {
		host = &envServerAddr
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		resolveHost = &envBaseURL
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		loggingLevel = &envLogLevel
	}
	if envFilename := os.Getenv("FILE_STORAGE_PATH"); envFilename != "" {
		filename = &envFilename
	}

	return &Options{
		Host:         *host,
		ResolveHost:  *resolveHost,
		LoggingLevel: *loggingLevel,
		Filename:     *filename,
	}
}

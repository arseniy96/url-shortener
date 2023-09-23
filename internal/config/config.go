// Package config – отвечает за конфигурацию приложения
// Конфигурировать приложение можно как флагами в командной строке, так и env переменными
// Инициализация конфига происходит при старте приложения.
package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"dario.cat/mergo"
	"github.com/caarlos0/env"
)

// Options – структура, которая хранит все настройки приложения.
type Options struct {
	// Host – адрес, на котором запустится веб-сервер.
	Host string `env:"SERVER_ADDRESS" json:"server_address"`
	// ResolveHost – адрес, который используется для резолва сокращённой ссылки.
	ResolveHost string `env:"BASE_URL" json:"base_url"`
	// LoggingLevel – уровень логирования.
	LoggingLevel string `env:"LOG_LEVEL" enums:"debug,info,warn,error" json:"log_level"`
	// Filename – путь к файлу, который будет выступать в качестве хранилища.
	Filename string `env:"FILE_STORAGE_PATH" json:"file_storage_path"`
	// ConnectionData – DSN для БД.
	ConnectionData string `env:"DATABASE_DSN" json:"database_dsn"`
	// EnableHTTPS – HTTPS mode
	EnableHTTPS bool `env:"ENABLE_HTTPS" json:"enable_https"`
	// ConfigPath – путь до файла с настройками
	ConfigPath string `env:"CONFIG"`
}

// InitConfig – функция для инициализации конфигурации приложения.
func InitConfig() *Options {
	options := &Options{}

	flag.StringVar(&options.Host, "a", "", "server host with port")
	flag.StringVar(&options.ResolveHost, "b", "", "resolve link address")
	flag.StringVar(&options.LoggingLevel, "l", "", "log level")
	flag.StringVar(&options.Filename, "f", "", "storage file")
	flag.StringVar(&options.ConnectionData, "d", "", "database connection data")
	flag.BoolVar(&options.EnableHTTPS, "s", false, "HTTPS mode")
	flag.StringVar(&options.ConfigPath, "c", "", "config path")
	flag.Parse()

	err := env.Parse(options)
	if err != nil {
		panic(err)
	}

	if options.ConfigPath != "" {
		_, err := mergeSettings(options, options.ConfigPath)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(options)

	return options
}

func mergeSettings(op *Options, configPath string) (*Options, error) {
	config := &Options{}
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(configFile).Decode(config)
	if err != nil {
		return nil, err
	}

	if err = mergo.Merge(op, config); err != nil {
		return nil, err
	}

	fmt.Println(config)
	fmt.Println(op)

	return op, err
}

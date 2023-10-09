// Package config – отвечает за конфигурацию приложения
// Конфигурировать приложение можно как флагами в командной строке, так и env переменными
// Инициализация конфига происходит при старте приложения.
package config

import (
	"encoding/json"
	"flag"
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
	// ConfigPath – путь до файла с настройками
	ConfigPath string `env:"CONFIG"`
	// TrustedSubnet – разрешённая подсеть
	TrustedSubnet string `env:"TRUSTED_SUBNET" json:"trusted_subnet"`
	// EnableHTTPS – HTTPS mode
	EnableHTTPS bool `env:"ENABLE_HTTPS" json:"enable_https"`
}

// InitConfig – функция для инициализации конфигурации приложения.
func InitConfig() (*Options, error) {
	options := &Options{}

	flag.StringVar(&options.Host, "a", "localhost:8080", "server host with port")
	flag.StringVar(&options.ResolveHost, "b", "http://localhost:8080", "resolve link address")
	flag.StringVar(&options.LoggingLevel, "l", "info", "log level")
	flag.StringVar(&options.Filename, "f", "", "storage file")
	flag.StringVar(&options.ConnectionData, "d", "", "database connection data")
	flag.BoolVar(&options.EnableHTTPS, "s", false, "HTTPS mode")
	flag.StringVar(&options.ConfigPath, "c", "", "config path")
	flag.StringVar(&options.TrustedSubnet, "t", "", "trusted subnet for GET stats")
	flag.Parse()

	err := env.Parse(options)
	if err != nil {
		return nil, err
	}

	if options.ConfigPath != "" {
		_, err := mergeSettings(options, options.ConfigPath)
		if err != nil {
			return nil, err
		}
	}

	return options, nil
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

	return op, err
}

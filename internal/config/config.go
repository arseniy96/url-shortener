// Package config – отвечает за конфигурацию приложения
// Конфигурировать приложение можно как флагами в командной строке, так и env переменными
// Инициализация конфига происходит при старте приложения
package config

import (
	"flag"

	"github.com/caarlos0/env"
)

// Options – структура, которая хранит все настройки приложения
type Options struct {
	// Host – адрес, на котором запустится веб-сервер
	Host string `env:"SERVER_ADDRESS"`
	// ResolveHost – адрес, который используется для резолва сокращённой ссылки
	ResolveHost string `env:"BASE_URL"`
	// LoggingLevel – уровень логирования
	LoggingLevel string `env:"LOG_LEVEL" enums:"debug,info,warn,error"`
	// Filename – путь к файлу, который будет выступать в качестве хранилища
	Filename string `env:"FILE_STORAGE_PATH"`
	// ConnectionData – DSN для БД
	ConnectionData string `env:"DATABASE_DSN"`
}

// InitConfig – функция для инициализации конфигурации приложения
func InitConfig() *Options {
	options := &Options{}

	flag.StringVar(&options.Host, "a", "localhost:8080", "server host with port")
	flag.StringVar(&options.ResolveHost, "b", "http://localhost:8080", "resolve link address")
	flag.StringVar(&options.LoggingLevel, "l", "info", "log level")
	flag.StringVar(&options.Filename, "f", "/tmp/short-url-db.json", "storage file")
	flag.StringVar(&options.ConnectionData, "d", "", "database connection data")
	flag.Parse()

	err := env.Parse(options)
	if err != nil {
		panic(err)
	}

	return options
}

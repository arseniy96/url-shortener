package main

import (
	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/router"
	"github.com/arseniy96/url-shortener/internal/storage"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	appConfig := config.InitConfig()

	if err := logger.Initialize(appConfig.LoggingLevel); err != nil {
		return err
	}

	serverStorage, err := storage.NewStorage(appConfig.Filename)
	if err != nil {
		return err
	}
	if appConfig.Filename != "" {
		if err := serverStorage.Restore(); err != nil {
			logger.Log.Error("Restore storage error", zap.Error(err))
		}
	}

	s := handlers.NewServer(serverStorage, appConfig)
	r := router.NewRouter(s)

	logger.Log.Infow("Running server", "address", s.Config.Host)
	return http.ListenAndServe(s.Config.Host, r)
}

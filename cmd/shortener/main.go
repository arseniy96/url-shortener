package main

import (
	"net/http"
	_ "net/http/pprof"

	"go.uber.org/zap"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/router"
	"github.com/arseniy96/url-shortener/internal/services/mycrypto"
	"github.com/arseniy96/url-shortener/internal/storage"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

// @Title URLShortener API
// @Description Сервис сокращения URL.
// @Version 1.0.
// @Contact.email arsenzhar@yandex.ru.
// @BasePath /
// @Host localhost:8080.

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

	serverStorage, err := storage.NewStorage(appConfig.Filename, appConfig.ConnectionData)
	if err != nil {
		return err
	}
	defer func() {
		if err := serverStorage.CloseConnection(); err != nil {
			logger.Log.Error(err)
		}
	}()

	if err := serverStorage.Restore(); err != nil {
		logger.Log.Error("Restore storage error", zap.Error(err))
	}

	s := handlers.NewServer(serverStorage, appConfig)
	r := router.NewRouter(s)

	logger.Log.Infof("Build version: %v", buildVersion)
	logger.Log.Infof("Build date: %v", buildDate)
	logger.Log.Infof("Build commit: %v", buildCommit)
	logger.Log.Infow("Running server", "address", s.Config.Host)
	if appConfig.EnableHTTPS {
		certFile, keyFile, err := mycrypto.LoadCryptoFiles()
		if err != nil {
			return err
		}
		return http.ListenAndServeTLS(s.Config.Host, certFile, keyFile, r)
	}
	return http.ListenAndServe(s.Config.Host, r)
}

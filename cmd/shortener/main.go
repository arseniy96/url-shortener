package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

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

	srv := http.Server{Addr: s.Config.Host, Handler: r}
	conClosed := make(chan struct{})
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-sigint
		// получили сигнал
		fmt.Println("get signal")
		if err := srv.Shutdown(context.Background()); err != nil {
			logger.Log.Errorf("HTTP server Shutdown: %v", err)
		}
		close(conClosed)
	}()

	var finalErr error
	if appConfig.EnableHTTPS {
		certFile, keyFile, err := mycrypto.LoadCryptoFiles()
		if err != nil {
			return err
		}
		finalErr = srv.ListenAndServeTLS(certFile, keyFile)
	} else {
		finalErr = srv.ListenAndServe()
	}

	<-conClosed
	fmt.Println("graceful shutdown")

	if errors.Is(finalErr, http.ErrServerClosed) {
		return nil
	}

	return finalErr
}

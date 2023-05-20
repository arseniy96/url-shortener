package main

import (
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/router"
	"github.com/arseniy96/url-shortener/internal/server"
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
	serverStorage := storage.NewStorage()
	s := server.NewServer(serverStorage)

	if err := logger.Initialize(s.Config.LoggingLevel); err != nil {
		return err
	}

	r := router.NewRouter(s)

	logger.Log.Info("Running server", zap.String("address", s.Config.Host))
	return http.ListenAndServe(s.Config.Host, r)
}

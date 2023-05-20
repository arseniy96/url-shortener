package main

import (
	"github.com/arseniy96/url-shortener/internal/router"
	"github.com/arseniy96/url-shortener/internal/server"
	"github.com/arseniy96/url-shortener/internal/storage"
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
	r := router.NewRouter(s)

	return http.ListenAndServe(s.Config.Host, r)
}

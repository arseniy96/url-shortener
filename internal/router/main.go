package router

import (
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/server"
	"github.com/go-chi/chi/v5"
)

func NewRouter(server server.Server) chi.Router {
	router := chi.NewRouter()
	router.Post("/", logger.RequestWithLog(server.CreateLink))
	router.Get("/{url_id}", logger.RequestWithLog(server.ResolveLink))

	return router
}

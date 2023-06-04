package router

import (
	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

func NewRouter(server *handlers.Server) chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares.GzipMiddleware, middlewares.LoggerMiddleware)
	router.Post("/", server.CreateLink)
	router.Get("/{url_id}", server.ResolveLink)
	router.Post("/api/shorten", server.CreateLinkJSON)
	router.Get("/ping", server.Ping)

	return router
}

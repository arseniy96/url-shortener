package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/middlewares"
)

// NewRouter – функция инициализации роутера
func NewRouter(server *handlers.Server) chi.Router {
	router := chi.NewRouter()
	router.Use(middlewares.GzipMiddleware, middlewares.LoggerMiddleware)
	router.Post("/", server.CookieMiddleware(server.CreateLink))
	router.Get("/{url_id}", server.CookieMiddleware(server.ResolveLink))
	router.Get("/ping", server.CookieMiddleware(server.Ping))
	router.Post("/api/shorten", server.CookieMiddleware(server.CreateLinkJSON))
	router.Post("/api/shorten/batch", server.CookieMiddleware(server.CreateLinksBatch))
	router.Get("/api/user/urls", server.CookieMiddleware(server.UserUrls))
	router.Delete("/api/user/urls", server.CookieMiddleware(server.DeleteUserUrls))

	router.Mount("/debug/", middleware.Profiler())

	return router
}

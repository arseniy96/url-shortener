package handlers

import (
	"context"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/services/keygenerator"
	"github.com/arseniy96/url-shortener/internal/storage"
)

type Repository interface {
	Add(string, string) error
	Get(string) (string, bool)
	HealthCheck() error
	GetMode() int
	AddBatch(context.Context, []storage.Record) error
	GetByOriginURL(string) (string, error)
}

type Generate interface {
	CreateKey() string
}

type Server struct {
	storage   Repository
	generator Generate
	Config    *config.Options
}

func NewServer(s Repository, c *config.Options) *Server {
	return &Server{
		storage:   s,
		generator: keygenerator.NewGenerator(s),
		Config:    c,
	}
}

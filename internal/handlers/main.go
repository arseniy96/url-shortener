package handlers

import (
	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/services/keygenerator"
	"github.com/arseniy96/url-shortener/internal/storage"
)

type Repository interface {
	Add(string, string)
	Get(string) (string, bool)
	HealthCheck() error
	GetMode() int
	AddBatch(records []storage.Record) error
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

package handlers

import (
	"context"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/services/keygenerator"
	"github.com/arseniy96/url-shortener/internal/storage"
)

type Repository interface {
	Add(string, string, string) error
	Get(string) (string, error)
	HealthCheck() error
	GetMode() int
	AddBatch(context.Context, []storage.Record) error
	GetByOriginURL(string) (string, error)
	GetByUser(context.Context, string) ([]storage.Record, error)
	CreateUser(context.Context) (*storage.User, error)
	UpdateUser(context.Context, int, string) error
	FindUserByID(context.Context, int) (*storage.User, error)
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

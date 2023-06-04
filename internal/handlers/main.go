package handlers

import (
	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/services/keygenerator"
)

type Repository interface {
	Add(string, string)
	Get(string) (string, bool)
}

type Generate interface {
	CreateKey() string
}

type database interface {
	HealthCheck() error
}

type Server struct {
	storage   Repository
	generator Generate
	Config    *config.Options
	database  database
}

func NewServer(s Repository, c *config.Options) *Server {
	return &Server{
		storage:   s,
		generator: keygenerator.NewGenerator(s),
		Config:    c,
	}
}

func (s *Server) SetDatabase(db database) {
	s.database = db
}

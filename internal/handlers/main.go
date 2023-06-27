package handlers

import (
	"context"
	"time"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/logger"
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
	DeleteUserURLs(context.Context, storage.DeleteURLMessage) error
}

type Generate interface {
	CreateKey() string
}

type Server struct {
	storage        Repository
	generator      Generate
	Config         *config.Options
	DeleteURLSChan chan storage.DeleteURLMessage
}

func NewServer(s Repository, c *config.Options) *Server {
	server := &Server{
		storage:        s,
		generator:      keygenerator.NewGenerator(s),
		Config:         c,
		DeleteURLSChan: make(chan storage.DeleteURLMessage, 10),
	}

	go server.deleteMessageBatch()

	return server
}

func (s *Server) deleteMessageBatch() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for { // FIXME: переделать на Fan-In
		select {
		case msg := <-s.DeleteURLSChan:
			err := s.storage.DeleteUserURLs(ctx, msg)
			if err != nil {
				logger.Log.Error(err)
				continue
			}
		default:
			continue
		}
	}
}

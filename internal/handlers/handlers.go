// Package handlers – содержит реализацию обработчиков запросов.
package handlers

import (
	"context"
	"fmt"
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
	DeleteUserURLs(storage.DeleteURLMessage) error
	GetURLsCount(context.Context) (int, error)
	GetUsersCount(context.Context) (int, error)
}

type Generate interface {
	CreateKey() string
}

type Server struct {
	Storage        Repository
	Generator      Generate
	Config         *config.Options
	DeleteURLSChan chan storage.DeleteURLMessage
}

const (
	CookieName             = "shortener_session"
	ContentTypeJSON        = "application/json"
	ContentTypeHeader      = "Content-Type"
	DeleteURLSChanSize     = 10
	InternalBackendErrTxt  = "Internal Backend Error"
	InvalidCookieErrTxt    = "Invalid Cookie"
	InvalidRequestErrTxt   = "Invalid request"
	UserUnauthorizedErrTxt = "User unauthorized"
	TimeOut                = 3 * time.Second
)

func NewServer(s Repository, c *config.Options) *Server {
	server := &Server{
		Storage:        s,
		Generator:      keygenerator.NewGenerator(s),
		Config:         c,
		DeleteURLSChan: make(chan storage.DeleteURLMessage, DeleteURLSChanSize),
	}

	go server.deleteMessageBatch()

	return server
}

func (s *Server) deleteMessageBatch() {
	for msg := range s.DeleteURLSChan {
		err := s.Storage.DeleteUserURLs(msg)
		if err != nil {
			logger.Log.Error(err)
			continue
		}
	}
}

func buildShortURL(host, path string) string {
	return fmt.Sprintf("%s/%s", host, path)
}

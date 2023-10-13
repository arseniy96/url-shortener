package grpchandlers

import (
	"context"
	"fmt"
	"time"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/handlers"
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

type GRPCServer struct {
	UnimplementedShortenerProtoServer
	handlers.Server
}

const (
	TimeOut = 3 * time.Second
)

func NewServer(s Repository, c *config.Options) *GRPCServer {
	server := &GRPCServer{
		Server: handlers.Server{
			Storage:   s,
			Generator: keygenerator.NewGenerator(s),
			Config:    c,
		},
	}

	return server
}

func buildShortURL(host, path string) string {
	return fmt.Sprintf("%s/%s", host, path)
}

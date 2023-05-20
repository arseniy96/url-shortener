package server

import (
	"fmt"
	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/keygenerator"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type Repository interface {
	Add(string, string)
	Get(string) (string, bool)
}

type Generate interface {
	CreateKey() string
}

type Server struct {
	storage   Repository
	generator Generate
	Config    *config.Options
}

func NewServer(s Repository) Server {
	return Server{
		storage:   s,
		generator: keygenerator.NewGenerator(s),
		Config:    config.SetConfig(),
	}
}

func (s Server) CreateLink(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil || len(body) == 0 {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	key := s.generator.CreateKey()
	s.storage.Add(key, string(body))
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("%s/%s", s.Config.ResolveHost, key)))
}

func (s Server) ResolveLink(writer http.ResponseWriter, request *http.Request) {
	urlID := chi.URLParam(request, "url_id")

	url, ok := s.storage.Get(urlID)
	if !ok {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Location", url)
	writer.WriteHeader(http.StatusTemporaryRedirect)
}

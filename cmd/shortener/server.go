package main

import (
	"fmt"
	"github.com/arseniy96/url-shortener/cmd/config"
	"github.com/arseniy96/url-shortener/cmd/storage"
	"github.com/arseniy96/url-shortener/cmd/utils"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type Server struct {
	storage   storage.Repository
	generator utils.Generate
	config    *config.Options
}

func NewServer(s storage.Repository) Server {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	return Server{
		storage:   s,
		generator: utils.NewGenerator(letters, s),
		config:    config.SetConfig(*host, *resolveHost),
	}
}

func (s Server) CreateLink(writer http.ResponseWriter, request *http.Request) {
	// check request
	// parse body
	// generate key
	// save url
	// response

	body, err := io.ReadAll(request.Body)
	if err != nil || len(body) == 0 {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	key := s.generator.CreateKey()
	s.storage.Add(key, string(body))
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("%s/%s", s.config.ResolveHost, key)))
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

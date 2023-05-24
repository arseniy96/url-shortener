package server

import (
	"encoding/json"
	"fmt"
	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/keygenerator"
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/models"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
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
	logger.Log.Debug(
		"Parsed body",
		zap.String("body", string(body)),
	)
	if err != nil {
		logger.Log.Error(
			"Invalid Request",
			zap.Error(err),
		)
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		logger.Log.Error(
			"Invalid body",
		)
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	key := s.generator.CreateKey()

	logger.Log.Debug(
		"Key was generated",
		zap.String("key", key),
	)

	s.storage.Add(key, string(body))
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("%s/%s", s.Config.ResolveHost, key)))
}

func (s Server) ResolveLink(writer http.ResponseWriter, request *http.Request) {
	urlID := chi.URLParam(request, "url_id")

	logger.Log.Debug(
		"URL was parsed",
		zap.String("url_id", urlID),
	)

	url, ok := s.storage.Get(urlID)
	if !ok {
		logger.Log.Error(
			"Unknown URL key",
			zap.String("key", urlID),
		)
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	logger.Log.Debug("Stored URL", zap.String("url", url))
	writer.Header().Set("Location", url)
	writer.WriteHeader(http.StatusTemporaryRedirect)
}

func (s Server) APICreateLink(writer http.ResponseWriter, request *http.Request) {
	var body models.RequestCreateLink
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&body); err != nil {
		logger.Log.Error(
			"Invalid Request",
			zap.Error(err),
		)
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}
	url := body.URL

	logger.Log.Debug(
		"Parsed URL",
		zap.String("URL", url),
	)

	if len(url) == 0 {
		logger.Log.Error(
			"Invalid URL",
		)
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	key := s.generator.CreateKey()
	s.storage.Add(key, url)

	logger.Log.Debug(
		"Key was generated",
		zap.String("key", key),
	)

	resp := models.ResponseCreateLink{
		Result: fmt.Sprintf("%s/%s", s.Config.ResolveHost, key),
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(resp); err != nil {
		logger.Log.Error(
			"error encoding response",
			zap.Error(err),
		)
		http.Error(writer, "Internal Backend Error", http.StatusBadRequest)
		return
	}
}

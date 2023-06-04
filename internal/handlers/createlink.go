package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/arseniy96/url-shortener/internal/models"
	"io"
	"net/http"
)

func (s *Server) CreateLink(writer http.ResponseWriter, request *http.Request) {
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

func (s *Server) CreateLinkJSON(writer http.ResponseWriter, request *http.Request) {
	var body models.RequestCreateLink
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&body); err != nil {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}
	url := body.URL

	if len(url) == 0 {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	key := s.generator.CreateKey()
	s.storage.Add(key, url)

	resp := models.ResponseCreateLink{
		Result: fmt.Sprintf("%s/%s", s.Config.ResolveHost, key),
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(resp); err != nil {
		http.Error(writer, "Internal Backend Error", http.StatusInternalServerError)
		return
	}
}

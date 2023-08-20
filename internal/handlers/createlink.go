package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/models"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func (s *Server) CreateLink(writer http.ResponseWriter, request *http.Request) {
	var resp []byte
	var respStatus int

	body, err := io.ReadAll(request.Body)
	if err != nil || len(body) == 0 {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}
	userSession, err := request.Cookie("shortener_session")
	if err != nil {
		http.Error(writer, "User unauthorized", http.StatusBadRequest)
		return
	}

	key := s.generator.CreateKey()

	err = s.storage.Add(key, string(body), userSession.Value)
	if err != nil {
		logger.Log.Error(err)
		if err == storage.ErrConflict {
			shortURL, err := s.storage.GetByOriginURL(string(body))
			if err != nil {
				http.Error(writer, "Internal Backend Error", http.StatusInternalServerError)
				return
			}
			respStatus = http.StatusConflict
			resp = []byte(fmt.Sprintf("%s/%s", s.Config.ResolveHost, shortURL))
		} else {
			http.Error(writer, "Internal Backend Error", http.StatusInternalServerError)
			return
		}
	} else {
		respStatus = http.StatusCreated
		resp = []byte(fmt.Sprintf("%s/%s", s.Config.ResolveHost, key))
	}
	writer.WriteHeader(respStatus)
	writer.Write(resp)
}

func (s *Server) CreateLinkJSON(writer http.ResponseWriter, request *http.Request) {
	var body models.RequestCreateLink
	var resp models.ResponseCreateLink
	var respStatus int

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&body); err != nil {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}
	url := body.URL

	userSession, err := request.Cookie("shortener_session")
	if err != nil {
		http.Error(writer, "User unauthorized", http.StatusBadRequest)
		return
	}

	if len(url) == 0 {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	key := s.generator.CreateKey()
	err = s.storage.Add(key, url, userSession.Value)
	if err != nil {
		if err == storage.ErrConflict {
			shortURL, err := s.storage.GetByOriginURL(url)
			if err != nil {
				http.Error(writer, "Internal Backend Error", http.StatusInternalServerError)
				return
			}
			respStatus = http.StatusConflict
			resp = models.ResponseCreateLink{
				Result: fmt.Sprintf("%s/%s", s.Config.ResolveHost, shortURL),
			}
		} else {
			http.Error(writer, "Internal Backend Error", http.StatusInternalServerError)
			return
		}
	} else {
		resp = models.ResponseCreateLink{
			Result: fmt.Sprintf("%s/%s", s.Config.ResolveHost, key),
		}
		respStatus = http.StatusCreated
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(respStatus)

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(resp); err != nil {
		http.Error(writer, "Internal Backend Error", http.StatusInternalServerError)
		return
	}
}

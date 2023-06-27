package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/models"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func (s *Server) DeleteUserUrls(writer http.ResponseWriter, request *http.Request) {
	var urls models.RequestDeleteUserURLS

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&urls); err != nil {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	userCookie, err := request.Cookie("shortener_session")
	if err != nil {
		logger.Log.Error(err)
		http.Error(writer, "Invalid cookie", http.StatusUnauthorized)
		return
	}

	s.DeleteURLSChan <- storage.DeleteURLMessage{
		ShortURLs:  urls,
		UserCookie: userCookie.Value,
	}

	writer.WriteHeader(http.StatusAccepted)
	writer.Write([]byte("ok"))
}

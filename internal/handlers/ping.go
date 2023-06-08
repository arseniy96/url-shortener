package handlers

import (
	"net/http"

	"github.com/arseniy96/url-shortener/internal/logger"
)

func (s *Server) Ping(writer http.ResponseWriter, request *http.Request) {
	if s.storage.GetMode() != 2 {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := s.storage.HealthCheck()
	if err != nil {
		logger.Log.Error(err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("OK"))
}

package handlers

import (
	"github.com/arseniy96/url-shortener/internal/logger"
	"net/http"
)

func (s *Server) Ping(writer http.ResponseWriter, request *http.Request) {
	if s.database == nil {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err := s.database.HealthCheck()
	if err != nil {
		logger.Log.Error(err)
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("OK"))
}

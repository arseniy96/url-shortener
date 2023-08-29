package handlers

import (
	"net/http"

	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func (s *Server) Ping(writer http.ResponseWriter, request *http.Request) {
	if s.storage.GetMode() != storage.DBMode {
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}

	err := s.storage.HealthCheck()
	if err != nil {
		logger.Log.Error(err)
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write([]byte("OK"))
	if err != nil {
		logger.Log.Error(err)
	}
}

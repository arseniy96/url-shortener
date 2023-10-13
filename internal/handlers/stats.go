package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/models"
)

func (s *Server) Stats(writer http.ResponseWriter, request *http.Request) {
	ctx, cancelCtx := context.WithTimeout(context.Background(), TimeOut)
	defer cancelCtx()
	urlsCount, err := s.Storage.GetURLsCount(ctx)
	if err != nil {
		logger.Log.Errorf("get urls count error: %v", err)
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}

	ctx, cancelCtx = context.WithTimeout(context.Background(), TimeOut)
	defer cancelCtx()
	usersCount, err := s.Storage.GetUsersCount(ctx)
	if err != nil {
		logger.Log.Errorf("get users count error: %v", err)
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}

	resp := models.ResponseStats{
		URLs:  urlsCount,
		Users: usersCount,
	}
	writer.Header().Set(ContentTypeHeader, ContentTypeJSON)
	writer.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(resp); err != nil {
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}
}

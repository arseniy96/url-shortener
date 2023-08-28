package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/models"
)

func (s *Server) UserUrls(writer http.ResponseWriter, request *http.Request) {
	userCookie, err := request.Cookie(CookieName)
	if err != nil {
		logger.Log.Error(err)
		http.Error(writer, InvalidCookieErrTxt, http.StatusUnauthorized)
		return
	}
	ctx, cancelCtx := context.WithTimeout(context.Background(), TimeOut)
	defer cancelCtx()
	records, err := s.storage.GetByUser(ctx, userCookie.Value)
	if err != nil {
		logger.Log.Error(err)
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}
	if len(records) == 0 {
		http.Error(writer, "User doesn't have urls", http.StatusNoContent)
		return
	}

	response := models.ResponseUserURLS{}
	for _, rec := range records {
		respEl := models.ResponseUserURL{
			ShortURL:    fmt.Sprintf("%s/%s", s.Config.ResolveHost, rec.ShortULR),
			OriginalURL: rec.OriginalURL,
		}
		response = append(response, respEl)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(response); err != nil {
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}
}

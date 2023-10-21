package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/arseniy96/url-shortener/internal/storage"
)

// ResolveLink godoc
// @Summary      Делает редирект на оригинальную ссылку
// @Description  По сокращённой ссылке делает редирект на оригинальную ссылку
// @Param		 url_id path string true "ID Short URL" example(maIJa1)
// @Success      307
// @Failure		 500 {object} object{} "Ошибка сервера"
// @Header		 307 {string} Location "http://localhost:8080/maIJa1"
// @Router       /{url_id} [get] .
func (s *Server) ResolveLink(writer http.ResponseWriter, request *http.Request) {
	urlID := chi.URLParam(request, "url_id")

	url, err := s.Storage.Get(urlID)
	if err != nil {
		if errors.Is(err, storage.ErrDeleted) {
			http.Error(writer, "URL was deleted", http.StatusGone)
			return
		}
		http.Error(writer, InvalidRequestErrTxt, http.StatusBadRequest)
		return
	}

	writer.Header().Set("Location", url)
	writer.WriteHeader(http.StatusTemporaryRedirect)
}

package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/arseniy96/url-shortener/internal/storage"
)

func (s *Server) ResolveLink(writer http.ResponseWriter, request *http.Request) {
	urlID := chi.URLParam(request, "url_id")

	url, err := s.storage.Get(urlID)
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

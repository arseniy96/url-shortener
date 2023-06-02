package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (s *Server) ResolveLink(writer http.ResponseWriter, request *http.Request) {
	urlID := chi.URLParam(request, "url_id")

	url, ok := s.storage.Get(urlID)
	if !ok {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Location", url)
	writer.WriteHeader(http.StatusTemporaryRedirect)
}

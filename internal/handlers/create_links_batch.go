package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/arseniy96/url-shortener/internal/models"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func (s *Server) CreateLinksBatch(writer http.ResponseWriter, request *http.Request) {
	var body models.RequestCreateLinksBatch
	var response models.ResponseCreateLinksBatch

	records := make([]storage.Record, 0)

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&body); err != nil {
		http.Error(writer, InvalidRequestErrTxt, http.StatusBadRequest)
		return
	}

	for _, el := range body {
		key := s.generator.CreateKey()

		rec := storage.Record{
			UUID:        el.CorrelationID,
			OriginalURL: el.OriginalURL,
			ShortULR:    key,
		}

		records = append(records, rec)

		response = append(response, models.ResponseLinks{
			CorrelationID: rec.UUID,
			ShortURL:      buildShortURL(s.Config.ResolveHost, key),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), TimeOut)
	defer cancel()
	err := s.storage.AddBatch(ctx, records)
	if err != nil {
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}

	writer.Header().Set(ContentTypeHeader, ContentTypeJSON)
	writer.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(response); err != nil {
		http.Error(writer, InternalBackendErrTxt, http.StatusInternalServerError)
		return
	}
}

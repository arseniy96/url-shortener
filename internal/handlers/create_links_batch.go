package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/arseniy96/url-shortener/internal/models"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func (s *Server) CreateLinksBatch(writer http.ResponseWriter, request *http.Request) {
	var body models.RequestCreateLinksBatch
	var records []storage.Record
	var response models.ResponseCreateLinksBatch

	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&body); err != nil {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
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
			ShortURL:      fmt.Sprintf("%s/%s", s.Config.ResolveHost, key),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := s.storage.AddBatch(ctx, records)
	if err != nil {
		http.Error(writer, "Internal Backend Error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(response); err != nil {
		http.Error(writer, "Internal Backend Error", http.StatusInternalServerError)
		return
	}
}

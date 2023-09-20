package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func TestServer_CreateLinksBatch(t *testing.T) {
	c := &config.Options{
		Host:         "localhost:8080",
		ResolveHost:  "http://localhost:8080",
		LoggingLevel: "debug",
	}

	type fields struct {
		storage   Repository
		generator Generate
		config    *config.Options
	}
	type want struct {
		expectedResponse string
		expectedCode     int
	}
	tests := []struct {
		name   string
		body   string
		fields fields
		want   want
	}{
		{
			name: "success result",
			body: `[{"correlation_id": "58039b0a-480d-11ee-9ace-0e6250a0eb02", "original_url": "https://ya.ru"}]`,
			fields: fields{
				storage:   NewTestStorage(storage.DBMode),
				generator: NewTestGenerator(),
				config:    c,
			},
			want: want{
				expectedResponse: "[{\"correlation_id\":\"58039b0a-480d-11ee-9ace-0e6250a0eb02\"," +
					"\"short_url\":\"http://localhost:8080/test\"}]\n",
				expectedCode: http.StatusCreated,
			},
		},
		{
			name: "invalid request",
			body: `https://ya.ru`,
			fields: fields{
				storage:   NewTestStorage(storage.DBMode),
				generator: NewTestGenerator(),
				config:    c,
			},
			want: want{
				expectedResponse: "Invalid request\n",
				expectedCode:     http.StatusBadRequest,
			},
		},
		{
			name: "success result",
			body: `[{"correlation_id": "58039b0a-480d-11ee-9ace-0e6250a0eb02", "original_url": "https://bad_test.ru"}]`,
			fields: fields{
				storage:   NewTestStorage(storage.DBMode),
				generator: NewTestGenerator(),
				config:    c,
			},
			want: want{
				expectedResponse: "Internal Backend Error\n",
				expectedCode:     http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/api/shorten/batch", strings.NewReader(tt.body))

			s := Server{
				storage:   tt.fields.storage,
				generator: tt.fields.generator,
				Config:    tt.fields.config,
			}

			s.CreateLinksBatch(writer, request)

			assert.Equal(t, tt.want.expectedCode, writer.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tt.want.expectedResponse, writer.Body.String(), "Тело ответа не совпадает с ожидаемым")
		})
	}
}

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func TestServer_Ping(t *testing.T) {
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
		fields fields
		want   want
	}{
		{
			name: "success result",
			fields: fields{
				storage:   NewTestStorage(storage.DBMode),
				generator: NewTestGenerator(),
				config:    c,
			},
			want: want{
				expectedResponse: "OK",
				expectedCode:     http.StatusOK,
			},
		},
		{
			name: "failed result",
			fields: fields{
				storage:   NewTestStorage(storage.MemoryMode),
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
			request := httptest.NewRequest(http.MethodPost, "/ping", nil)

			s := Server{
				Storage:   tt.fields.storage,
				Generator: tt.fields.generator,
				Config:    tt.fields.config,
			}

			s.Ping(writer, request)

			assert.Equal(t, tt.want.expectedCode, writer.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tt.want.expectedResponse, writer.Body.String(), "Тело ответа не совпадает с ожидаемым")
		})
	}
}

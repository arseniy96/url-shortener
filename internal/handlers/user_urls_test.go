package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func TestServer_UserUrls(t *testing.T) {
	c := &config.Options{
		Host:         "localhost:8080",
		ResolveHost:  "http://localhost:8080",
		LoggingLevel: "debug",
	}

	if err := logger.Initialize(c.LoggingLevel); err != nil {
		panic(err)
	}

	type fields struct {
		storage     Repository
		generator   Generate
		config      *config.Options
		cookieName  string
		cookieValue string
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
			name: "Success result",
			fields: fields{
				storage:     NewTestStorage(storage.DBMode),
				generator:   NewTestGenerator(),
				config:      c,
				cookieName:  CookieName,
				cookieValue: "valid",
			},
			want: want{
				expectedResponse: "[{\"short_url\":\"http://localhost:8080/test\",\"original_url\":\"https://ya.ru\"}]\n",
				expectedCode:     http.StatusOK,
			},
		},
		{
			name: "Unauthorized",
			fields: fields{
				storage:     NewTestStorage(storage.DBMode),
				generator:   NewTestGenerator(),
				config:      c,
				cookieName:  "test",
				cookieValue: "valid",
			},
			want: want{
				expectedResponse: "Invalid Cookie\n",
				expectedCode:     http.StatusUnauthorized,
			},
		},
		{
			name: "not db mode",
			fields: fields{
				storage:     NewTestStorage(storage.DBMode),
				generator:   NewTestGenerator(),
				config:      c,
				cookieName:  CookieName,
				cookieValue: "invalid",
			},
			want: want{
				expectedResponse: "Internal Backend Error\n",
				expectedCode:     http.StatusInternalServerError,
			},
		},
		{
			name: "no content",
			fields: fields{
				storage:     NewTestStorage(storage.DBMode),
				generator:   NewTestGenerator(),
				config:      c,
				cookieName:  CookieName,
				cookieValue: "empty",
			},
			want: want{
				expectedResponse: "User doesn't have urls\n",
				expectedCode:     http.StatusNoContent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/api/user/urls", nil)

			s := Server{
				storage:   tt.fields.storage,
				generator: tt.fields.generator,
				Config:    tt.fields.config,
			}

			request.AddCookie(&http.Cookie{
				Name:  tt.fields.cookieName,
				Value: tt.fields.cookieValue,
			})

			s.UserUrls(writer, request)

			assert.Equal(t, tt.want.expectedCode, writer.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tt.want.expectedResponse, writer.Body.String(), "Тело ответа не совпадает с ожидаемым")
		})
	}
}

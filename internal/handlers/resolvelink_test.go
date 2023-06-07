package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/arseniy96/url-shortener/internal/config"
)

func TestServer_ResolveLink(t *testing.T) {
	c := &config.Options{
		Host:         "localhost:8080",
		ResolveHost:  "http://localhost:8080",
		LoggingLevel: "debug",
		Filename:     "",
	}
	type fields struct {
		storage   Repository
		generator Generate
		config    *config.Options
	}
	type want struct {
		expectedCode int
	}
	tests := []struct {
		name   string
		urlID  string
		fields fields
		want   want
	}{
		{
			name:  "Invalid request",
			urlID: "Some",
			fields: fields{
				storage:   NewTestStorage(),
				generator: NewTestGenerator(),
				config:    c,
			},
			want: want{
				expectedCode: 400,
			},
		},
		{
			name:  "Valid request",
			urlID: "test",
			fields: fields{
				storage:   NewTestStorage(),
				generator: NewTestGenerator(),
				config:    c,
			},
			want: want{
				expectedCode: 307,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, `/{url_id}`, nil)

			ctx := chi.NewRouteContext()
			ctx.URLParams.Add("url_id", tt.urlID)
			request = request.WithContext(context.WithValue(request.Context(), chi.RouteCtxKey, ctx))

			s := Server{
				storage:   tt.fields.storage,
				generator: tt.fields.generator,
				Config:    tt.fields.config,
			}

			s.ResolveLink(writer, request)

			assert.Equal(t, tt.want.expectedCode, writer.Code, "Код ответа не совпадает с ожидаемым")
		})
	}
}

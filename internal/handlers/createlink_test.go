package handlers

import (
	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer_CreateLink(t *testing.T) {
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
			name: "Invalid Body",
			body: "",
			fields: fields{
				storage:   NewTestStorage(),
				generator: NewTestGenerator(),
				config:    c,
			},
			want: want{
				expectedCode:     http.StatusBadRequest,
				expectedResponse: "Invalid request\n",
			},
		},
		{
			name: "Valid request",
			body: "https://ya.ru",
			fields: fields{
				storage:   NewTestStorage(),
				generator: NewTestGenerator(),
				config:    c,
			},
			want: want{
				expectedCode:     http.StatusCreated,
				expectedResponse: `http://localhost:8080/test`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))

			s := Server{
				storage:   tt.fields.storage,
				generator: tt.fields.generator,
				Config:    tt.fields.config,
			}

			s.CreateLink(writer, request)

			assert.Equal(t, tt.want.expectedCode, writer.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tt.want.expectedResponse, writer.Body.String(), "Тело ответа не совпадает с ожидаемым")
		})
	}
}

func TestServer_CreateLinkJSON(t *testing.T) {
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
			name: "Invalid Body",
			body: "{}",
			fields: fields{
				storage:   NewTestStorage(),
				generator: NewTestGenerator(),
				config:    c,
			},
			want: want{
				expectedCode:     http.StatusBadRequest,
				expectedResponse: "Invalid request\n",
			},
		},
		{
			name: "Valid request",
			body: `{ "url": "https://ya.ru" }`,
			fields: fields{
				storage:   NewTestStorage(),
				generator: NewTestGenerator(),
				config:    c,
			},
			want: want{
				expectedCode:     http.StatusCreated,
				expectedResponse: "{\"result\":\"http://localhost:8080/test\"}\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodPost, "/api/shorten", strings.NewReader(tt.body))

			s := Server{
				storage:   tt.fields.storage,
				generator: tt.fields.generator,
				Config:    tt.fields.config,
			}

			s.CreateLinkJSON(writer, request)

			assert.Equal(t, tt.want.expectedCode, writer.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tt.want.expectedResponse, writer.Body.String(), "Тело ответа не совпадает с ожидаемым")
		})
	}
}

type TestStorage struct {
	Urls map[string]string
}

func NewTestStorage() *TestStorage {
	return &TestStorage{
		Urls: make(map[string]string),
	}
}

func (s *TestStorage) Add(_, _ string) {
	s.Urls["test"] = "Test"
}

func (s *TestStorage) Get(key string) (string, bool) {
	if key == "test" {
		return "test", true
	} else {
		return "", false
	}
}

func (s *TestStorage) HealthCheck() error {
	return nil
}

type TestGenerator struct{}

func NewTestGenerator() *TestGenerator {
	return &TestGenerator{}
}

func (s *TestGenerator) CreateKey() string {
	return "test"
}

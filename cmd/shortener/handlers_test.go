package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

func TestGetHandler(t *testing.T) {
	tests := []struct {
		name         string
		method       string
		path         string
		expectedCode int
	}{
		{
			name:         "Invalid Request",
			method:       http.MethodGet,
			path:         "/",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid Request",
			method:       http.MethodGet,
			path:         "/Some",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Valid Request",
			method:       http.MethodGet,
			path:         "/test",
			expectedCode: http.StatusTemporaryRedirect,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()
			storage := NewTestStorage()

			GetHandler(w, r, storage)

			assert.Equal(t, tt.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
		})
	}
}

func TestMainHandlerNegative(t *testing.T) {
	tests := []struct {
		method       string
		expectedCode int
		expectedBody string
	}{
		{
			method:       http.MethodDelete,
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid request\n",
		},
		{
			method:       http.MethodPut,
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid request\n",
		},
		{
			method:       http.MethodHead,
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid request\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, "/", nil)
			w := httptest.NewRecorder()
			storage := NewTestStorage()

			MainHandler(storage)(w, r)

			assert.Equal(t, tt.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			// проверим корректность полученного тела ответа, если мы его ожидаем
			if tt.expectedBody != "" {
				// assert.JSONEq помогает сравнить две JSON-строки
				assert.Equal(t, tt.expectedBody, w.Body.String(), "Тело ответа не совпадает с ожидаемым")
			}
		})
	}
}

func TestPostHandler(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		path               string
		body               string
		expectedCode       int
		expectedBodyRegexp string
	}{
		{
			name:               "invalid path",
			method:             http.MethodGet,
			path:               "/some",
			body:               "http://ya.ru",
			expectedBodyRegexp: `Invalid path\n`,
			expectedCode:       http.StatusBadRequest,
		},
		{
			name:               "invalid body",
			method:             http.MethodGet,
			path:               "/",
			body:               "",
			expectedBodyRegexp: `Invalid request\n`,
			expectedCode:       http.StatusBadRequest,
		},
		{
			name:               "valid request",
			method:             http.MethodGet,
			path:               "/",
			body:               "http://ya.ru",
			expectedBodyRegexp: `^http\:\/\/localhost\:\d{1,6}\/\w{6}$`,
			expectedCode:       http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			w := httptest.NewRecorder()
			storage := NewTestStorage()

			PostHandler(w, r, storage)

			assert.Equal(t, tt.expectedCode, w.Code, "Код ответа не совпадает с ожидаемым")
			assert.Regexp(t, regexp.MustCompile(tt.expectedBodyRegexp), w.Body.String(), "Тело ответа не совпадает с ожидаемым")
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
		return "Test", true
	} else {
		return "", false
	}
}

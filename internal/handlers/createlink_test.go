package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/storage"
)

//nolint:dupl // it's ok
func TestServer_CreateLink(t *testing.T) {
	c := &config.Options{
		Host:         "localhost:8080",
		ResolveHost:  "http://localhost:8080",
		LoggingLevel: "debug",
	}
	if err := logger.Initialize(c.LoggingLevel); err != nil {
		panic(err)
	}

	type fields struct {
		storage    Repository
		generator  Generate
		config     *config.Options
		cookieName string
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
				storage:    NewTestStorage(storage.DBMode),
				generator:  NewTestGenerator(),
				config:     c,
				cookieName: CookieName,
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
				storage:    NewTestStorage(storage.DBMode),
				generator:  NewTestGenerator(),
				config:     c,
				cookieName: CookieName,
			},
			want: want{
				expectedCode:     http.StatusCreated,
				expectedResponse: `http://localhost:8080/test`,
			},
		},
		{
			name: "Unauthorized",
			body: "https://ya.ru",
			fields: fields{
				storage:    NewTestStorage(storage.DBMode),
				generator:  NewTestGenerator(),
				config:     c,
				cookieName: "InvalidCookie",
			},
			want: want{
				expectedCode:     http.StatusBadRequest,
				expectedResponse: "User unauthorized\n",
			},
		},
		{
			name: "duplicate error",
			body: "https://double_value.ru",
			fields: fields{
				storage:    NewTestStorage(storage.DBMode),
				generator:  NewTestGenerator(),
				config:     c,
				cookieName: CookieName,
			},
			want: want{
				expectedCode:     http.StatusConflict,
				expectedResponse: "http://localhost:8080/",
			},
		},
		{
			name: "internal error",
			body: "https://bad_test.ru",
			fields: fields{
				storage:    NewTestStorage(storage.DBMode),
				generator:  NewTestGenerator(),
				config:     c,
				cookieName: CookieName,
			},
			want: want{
				expectedCode:     http.StatusInternalServerError,
				expectedResponse: "Internal Backend Error\n",
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

			request.AddCookie(&http.Cookie{
				Name:  tt.fields.cookieName,
				Value: "test",
			})
			s.CreateLink(writer, request)

			assert.Equal(t, tt.want.expectedCode, writer.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tt.want.expectedResponse, writer.Body.String(), "Тело ответа не совпадает с ожидаемым")
		})
	}
}

//nolint:dupl // it's ok
func TestServer_CreateLinkJSON(t *testing.T) {
	c := &config.Options{
		Host:         "localhost:8080",
		ResolveHost:  "http://localhost:8080",
		LoggingLevel: "debug",
	}
	if err := logger.Initialize(c.LoggingLevel); err != nil {
		panic(err)
	}

	type fields struct {
		storage    Repository
		generator  Generate
		config     *config.Options
		cookieName string
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
				storage:    NewTestStorage(storage.DBMode),
				generator:  NewTestGenerator(),
				config:     c,
				cookieName: CookieName,
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
				storage:    NewTestStorage(storage.DBMode),
				generator:  NewTestGenerator(),
				config:     c,
				cookieName: CookieName,
			},
			want: want{
				expectedCode:     http.StatusCreated,
				expectedResponse: "{\"result\":\"http://localhost:8080/test\"}\n",
			},
		},
		{
			name: "Unauthorized",
			body: `{}`,
			fields: fields{
				storage:    NewTestStorage(storage.DBMode),
				generator:  NewTestGenerator(),
				config:     c,
				cookieName: "invalid",
			},
			want: want{
				expectedCode:     http.StatusBadRequest,
				expectedResponse: "User unauthorized\n",
			},
		},
		{
			name: "internal error",
			body: `{ "url": "https://bad_test.ru" }`,
			fields: fields{
				storage:    NewTestStorage(storage.DBMode),
				generator:  NewTestGenerator(),
				config:     c,
				cookieName: CookieName,
			},
			want: want{
				expectedCode:     http.StatusInternalServerError,
				expectedResponse: "Internal Backend Error\n",
			},
		},
		{
			name: "duplicate error",
			body: `{ "url": "https://double_value.ru" }`,
			fields: fields{
				storage:    NewTestStorage(storage.DBMode),
				generator:  NewTestGenerator(),
				config:     c,
				cookieName: CookieName,
			},
			want: want{
				expectedCode:     http.StatusConflict,
				expectedResponse: "{\"result\":\"http://localhost:8080/\"}\n",
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

			request.AddCookie(&http.Cookie{
				Name:  tt.fields.cookieName,
				Value: "test",
			})
			s.CreateLinkJSON(writer, request)

			assert.Equal(t, tt.want.expectedCode, writer.Code, "Код ответа не совпадает с ожидаемым")
			assert.Equal(t, tt.want.expectedResponse, writer.Body.String(), "Тело ответа не совпадает с ожидаемым")
		})
	}
}

type TestStorage struct {
	Urls map[string]string
	mode int
}

func NewTestStorage(mode int) *TestStorage {
	return &TestStorage{
		Urls: make(map[string]string),
		mode: mode,
	}
}

func (s *TestStorage) Add(_, value, _ string) error {
	switch value {
	case "https://bad_test.ru":
		return errors.New("SomeError")
	case "https://double_value.ru":
		return storage.ErrConflict
	default:
		s.Urls["test"] = "Test"
		return nil
	}
}

func (s *TestStorage) AddBatch(ctx context.Context, records []storage.Record) error {
	for _, rec := range records {
		if rec.OriginalURL == "https://bad_test.ru" {
			return fmt.Errorf("AddBatchError")
		}
	}
	return nil
}

func (s *TestStorage) Get(key string) (string, error) {
	switch key {
	case "test":
		return "test", nil
	case "deleted":
		return "", storage.ErrDeleted
	default:
		return "", fmt.Errorf("Error")
	}
}

func (s *TestStorage) GetByOriginURL(_ string) (string, error) {
	return "", nil
}

func (s *TestStorage) GetByUser(_ context.Context, cookie string) ([]storage.Record, error) {
	switch cookie {
	case "valid":
		return []storage.Record{{
			ShortULR:    "test",
			OriginalURL: "https://ya.ru",
		}}, nil
	case "invalid":
		return []storage.Record{}, storage.ErrNoDBMode
	case "empty":
		return []storage.Record{}, nil
	}
	return []storage.Record{}, nil
}

func (s *TestStorage) CreateUser(_ context.Context) (*storage.User, error) {
	var u *storage.User
	return u, nil
}

func (s *TestStorage) FindUserByID(_ context.Context, _ int) (*storage.User, error) {
	var u *storage.User
	return u, nil
}

func (s *TestStorage) DeleteUserURLs(storage.DeleteURLMessage) error {
	return nil
}

func (s *TestStorage) UpdateUser(_ context.Context, _ int, _ string) error {
	return nil
}

func (s *TestStorage) HealthCheck() error {
	return nil
}

func (s *TestStorage) GetMode() int {
	return s.mode
}

type TestGenerator struct{}

func NewTestGenerator() *TestGenerator {
	return &TestGenerator{}
}

func (s *TestGenerator) CreateKey() string {
	return "test"
}

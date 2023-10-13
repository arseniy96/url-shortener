package grpchandlers

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func TestGRPCServer_Ping(t *testing.T) {
	type fields struct {
		Server handlers.Server
	}
	tests := []struct {
		name    string
		fields  fields
		want    *PingResponse
		wantErr bool
	}{
		{
			name: "success pong",
			fields: fields{
				Server: handlers.Server{
					Storage: NewTestStorage(storage.DBMode),
				},
			},
			want:    &PingResponse{Result: "OK"},
			wantErr: false,
		},
		{
			name: "not db mode",
			fields: fields{
				Server: handlers.Server{
					Storage: NewTestStorage(storage.MemoryMode),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GRPCServer{
				Server: tt.fields.Server,
			}
			got, err := s.Ping(context.Background(), &PingRequest{})
			if (err != nil) != tt.wantErr {
				t.Errorf("Ping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ping() got = %v, want %v", got, tt.want)
			}
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

func (s *TestStorage) GetByOriginURL(key string) (string, error) {
	if key == "https://double_value.ru" {
		return "double_value", nil
	}
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

func (s *TestStorage) GetUsersCount(ctx context.Context) (int, error) {
	return 1, nil
}

func (s *TestStorage) GetURLsCount(ctx context.Context) (int, error) {
	return 1, nil
}

type TestGenerator struct{}

func NewTestGenerator() *TestGenerator {
	return &TestGenerator{}
}

func (s *TestGenerator) CreateKey() string {
	return "test"
}

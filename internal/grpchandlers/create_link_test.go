package grpchandlers

import (
	"context"
	"reflect"
	"testing"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func TestGRPCServer_CreateLink(t *testing.T) {
	if err := logger.Initialize("info"); err != nil {
		panic(err)
	}

	type fields struct {
		Server handlers.Server
	}
	type args struct {
		in *CreateLinkRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateLinkResponse
		wantErr bool
	}{
		{
			name: "success create link",
			fields: fields{
				Server: handlers.Server{
					Storage:   NewTestStorage(storage.MemoryMode),
					Generator: NewTestGenerator(),
					Config:    &config.Options{ResolveHost: "http://localhost:8080"},
				},
			},
			args: args{
				in: &CreateLinkRequest{
					Url:         "https://ya.ru",
					UserSession: "some_user",
				},
			},
			want: &CreateLinkResponse{
				Result: "http://localhost:8080/test",
			},
			wantErr: false,
		},
		{
			name: "duplicate link",
			fields: fields{
				Server: handlers.Server{
					Storage:   NewTestStorage(storage.DBMode),
					Generator: NewTestGenerator(),
					Config:    &config.Options{ResolveHost: "http://localhost:8080"},
				},
			},
			args: args{
				in: &CreateLinkRequest{
					Url:         "https://double_value.ru",
					UserSession: "some_user",
				},
			},
			want: &CreateLinkResponse{
				Result: "http://localhost:8080/double_value",
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				Server: handlers.Server{
					Storage:   NewTestStorage(storage.DBMode),
					Generator: NewTestGenerator(),
					Config:    &config.Options{ResolveHost: "http://localhost:8080"},
				},
			},
			args: args{
				in: &CreateLinkRequest{
					Url:         "https://bad_test.ru",
					UserSession: "some_user",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unauthorized",
			fields: fields{
				Server: handlers.Server{
					Storage:   NewTestStorage(storage.DBMode),
					Generator: NewTestGenerator(),
					Config:    &config.Options{ResolveHost: "http://localhost:8080"},
				},
			},
			args: args{
				in: &CreateLinkRequest{
					Url: "https://bad_test.ru",
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
			got, err := s.CreateLink(context.Background(), tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGRPCServer_CreateLinksBatch(t *testing.T) {
	if err := logger.Initialize("info"); err != nil {
		panic(err)
	}

	type fields struct {
		Server handlers.Server
	}
	type args struct {
		in *CreateLinksBatchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *CreateLinksBatchResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Server: handlers.Server{
					Storage:   NewTestStorage(storage.MemoryMode),
					Generator: NewTestGenerator(),
					Config:    &config.Options{ResolveHost: "http://localhost:8080"},
				},
			},
			args: args{
				in: &CreateLinksBatchRequest{
					Urls: []*CreateLinksBatchRequestNested{{
						CorrelationId: "some_id",
						OriginalUrl:   "https://yandex.ru",
					}},
					UserSession: "some_user",
				},
			},
			want: &CreateLinksBatchResponse{
				Urls: []*CreateLinksBatchResponseNested{{
					CorrelationId: "some_id",
					ShortUrl:      "http://localhost:8080/test",
				}},
			},
			wantErr: false,
		},
		{
			name: "unauthorized",
			fields: fields{
				Server: handlers.Server{
					Storage:   NewTestStorage(storage.MemoryMode),
					Generator: NewTestGenerator(),
					Config:    &config.Options{ResolveHost: "http://localhost:8080"},
				},
			},
			args: args{
				in: &CreateLinksBatchRequest{
					Urls: []*CreateLinksBatchRequestNested{{
						CorrelationId: "some_id",
						OriginalUrl:   "https://yandex.ru",
					}},
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
			got, err := s.CreateLinksBatch(context.Background(), tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateLinksBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateLinksBatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

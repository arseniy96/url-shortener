package grpchandlers

import (
	"context"
	"reflect"
	"testing"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/storage"
)

func TestGRPCServer_ResolveLink(t *testing.T) {
	type fields struct {
		Server handlers.Server
	}
	type args struct {
		in *ResolveLinkRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ResolveLinkResponse
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
				in: &ResolveLinkRequest{
					ShortUrl: "test",
				},
			},
			want: &ResolveLinkResponse{
				OriginalUrl: "test",
			},
			wantErr: false,
		},
		{
			name: "failed",
			fields: fields{
				Server: handlers.Server{
					Storage:   NewTestStorage(storage.MemoryMode),
					Generator: NewTestGenerator(),
					Config:    &config.Options{ResolveHost: "http://localhost:8080"},
				},
			},
			args: args{
				in: &ResolveLinkRequest{
					ShortUrl: "",
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
			got, err := s.ResolveLink(context.Background(), tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveLink() got = %v, want %v", got, tt.want)
			}
		})
	}
}

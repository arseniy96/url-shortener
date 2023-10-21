package grpchandlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/url-shortener/internal/storage"
)

func (s *GRPCServer) Ping(ctx context.Context, _ *PingRequest) (*PingResponse, error) {
	if s.Storage.GetMode() != storage.DBMode {
		return nil, status.Errorf(codes.Internal, "not DB mode")
	}

	err := s.Storage.HealthCheck()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "DB is broken")
	}

	resp := &PingResponse{
		Result: "OK",
	}

	return resp, nil
}

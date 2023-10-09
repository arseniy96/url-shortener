package grpchandlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/url-shortener/internal/storage"
	pb "github.com/arseniy96/url-shortener/src/proto"
)

func (s *GRPCServer) Ping(ctx context.Context, _ *pb.PingRequest) (*pb.PingResponse, error) {
	if s.Storage.GetMode() != storage.DBMode {
		return nil, status.Errorf(codes.Internal, "not DB mode")
	}

	err := s.Storage.HealthCheck()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "DB is broken")
	}

	resp := &pb.PingResponse{
		Result: "OK",
	}

	return resp, nil
}

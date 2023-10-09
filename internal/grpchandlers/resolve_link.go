package grpchandlers

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/storage"
	pb "github.com/arseniy96/url-shortener/src/proto"
)

func (s *GRPCServer) ResolveLink(ctx context.Context, in *pb.ResolveLinkRequest) (*pb.ResolveLinkResponse, error) {
	url, err := s.Storage.Get(in.GetShortUrl())
	if err != nil {
		if errors.Is(err, storage.ErrDeleted) {
			return nil, status.Errorf(codes.DataLoss, "URL was deleted")
		}
		return nil, status.Errorf(codes.DataLoss, handlers.InternalBackendErrTxt)
	}
	return &pb.ResolveLinkResponse{
		OriginalUrl: url,
	}, nil
}

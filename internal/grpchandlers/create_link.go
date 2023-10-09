package grpchandlers

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/storage"
	pb "github.com/arseniy96/url-shortener/src/proto"
)

func (s *GRPCServer) CreateLink(ctx context.Context, in *pb.CreateLinkRequest) (*pb.CreateLinkResponse, error) {
	originURL := in.Url
	userSession := in.UserSession
	if len(originURL) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "url length is zero")
	}
	if len(userSession) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, handlers.UserUnauthorizedErrTxt)
	}

	key := s.Generator.CreateKey()

	err := s.Storage.Add(key, originURL, userSession)
	if err != nil {
		logger.Log.Error(err)
		if errors.Is(err, storage.ErrConflict) {
			shortURL, err := s.Storage.GetByOriginURL(originURL)
			if err != nil {
				return nil, status.Errorf(codes.Internal, handlers.InternalBackendErrTxt)
			}
			resp := &pb.CreateLinkResponse{
				Result: buildShortURL(s.Config.ResolveHost, shortURL),
			}
			return resp, nil
		} else {
			return nil, status.Errorf(codes.Internal, handlers.InternalBackendErrTxt)
		}
	}
	resp := &pb.CreateLinkResponse{
		Result: buildShortURL(s.Config.ResolveHost, key),
	}
	return resp, nil
}

func (s *GRPCServer) CreateLinksBatch(ctx context.Context,
	in *pb.CreateLinksBatchRequest) (*pb.CreateLinksBatchResponse, error) {
	userSession := in.UserSession
	if len(userSession) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, handlers.UserUnauthorizedErrTxt)
	}

	response := make([]*pb.CreateLinksBatchResponseNested, 0)
	records := make([]storage.Record, 0)

	for _, url := range in.Urls {
		key := s.Generator.CreateKey()

		rec := storage.Record{
			UUID:        url.GetCorrelationId(),
			OriginalURL: url.GetOriginalUrl(),
			ShortULR:    key,
		}

		records = append(records, rec)

		response = append(response, &pb.CreateLinksBatchResponseNested{
			CorrelationId: rec.UUID,
			ShortUrl:      buildShortURL(s.Config.ResolveHost, key),
		})
	}

	err := s.Storage.AddBatch(ctx, records)
	if err != nil {
		return nil, status.Errorf(codes.Internal, handlers.InternalBackendErrTxt)
	}

	return &pb.CreateLinksBatchResponse{Urls: response}, nil
}

package grpchandlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/logger"
	pb "github.com/arseniy96/url-shortener/src/proto"
)

func (s *GRPCServer) UserUrls(ctx context.Context, in *pb.UserUrlsRequest) (*pb.UserUrlsResponse, error) {
	userSession := in.GetUserSession()
	if len(userSession) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, handlers.UserUnauthorizedErrTxt)
	}

	ctx2, cancelCtx := context.WithTimeout(context.Background(), TimeOut)
	defer cancelCtx()

	records, err := s.Storage.GetByUser(ctx2, userSession)
	if err != nil {
		logger.Log.Error(err)
		return nil, status.Errorf(codes.Internal, handlers.InternalBackendErrTxt)
	}
	if len(records) == 0 {
		return nil, status.Errorf(codes.NotFound, "User doesn't have urls")
	}

	urls := make([]*pb.UserUrlsResponseNested, 0)
	for _, rec := range records {
		respEl := &pb.UserUrlsResponseNested{
			ShortUrl:    buildShortURL(s.Config.ResolveHost, rec.ShortULR),
			OriginalUrl: rec.OriginalURL,
		}
		urls = append(urls, respEl)
	}

	return &pb.UserUrlsResponse{Urls: urls}, nil
}

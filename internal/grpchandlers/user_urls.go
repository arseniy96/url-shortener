package grpchandlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/logger"
)

func (s *GRPCServer) UserUrls(ctx context.Context, in *UserUrlsRequest) (*UserUrlsResponse, error) {
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

	urls := make([]*UserUrlsResponseNested, 0)
	for _, rec := range records {
		respEl := &UserUrlsResponseNested{
			ShortUrl:    buildShortURL(s.Config.ResolveHost, rec.ShortULR),
			OriginalUrl: rec.OriginalURL,
		}
		urls = append(urls, respEl)
	}

	return &UserUrlsResponse{Urls: urls}, nil
}

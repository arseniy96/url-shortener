package main

import (
	"context"
	"fmt"
	"net"

	"github.com/arseniy96/url-shortener/internal/handlers"
	pb "github.com/arseniy96/url-shortener/src/proto"

	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		return err
	}
	gRPCServer := grpc.NewServer()
	// регистрируем сервис
	pb.RegisterShortenerProtoServer(gRPCServer, &ShortenerServer{})

	fmt.Println("gRPC server is running")
	// получаем запрос gRPC
	return gRPCServer.Serve(listen)
}

type ShortenerServer struct {
	// нужно встраивать тип pb.Unimplemented<TypeName>
	// для совместимости с будущими версиями
	pb.UnimplementedShortenerProtoServer
	handlers.Server
}

func (s *ShortenerServer) Ping(ctx context.Context, request *pb.PingRequest) (*pb.PingResponse, error) {
	resp := &pb.PingResponse{
		Result: "OK",
	}

	return resp, nil
}

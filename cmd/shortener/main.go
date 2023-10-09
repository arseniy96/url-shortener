package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/arseniy96/url-shortener/internal/config"
	"github.com/arseniy96/url-shortener/internal/grpchandlers"
	"github.com/arseniy96/url-shortener/internal/handlers"
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/arseniy96/url-shortener/internal/router"
	"github.com/arseniy96/url-shortener/internal/services/mycrypto"
	"github.com/arseniy96/url-shortener/internal/storage"
	pb "github.com/arseniy96/url-shortener/src/proto"
)

const (
	timeoutServerShutdown = 5 * time.Second
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

// @Title URLShortener API
// @Description Сервис сокращения URL.
// @Version 1.0.
// @Contact.email arsenzhar@yandex.ru.
// @BasePath /
// @Host localhost:8080.

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	appConfig, err := config.InitConfig()
	if err != nil {
		return err
	}

	if err := logger.Initialize(appConfig.LoggingLevel); err != nil {
		return err
	}

	serverStorage, err := storage.NewStorage(appConfig.Filename, appConfig.ConnectionData)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	defer func() {
		if err := serverStorage.CloseConnection(); err != nil {
			logger.Log.Error(err)
		}
	}()

	if err := serverStorage.Restore(); err != nil {
		logger.Log.Error("Restore storage error", zap.Error(err))
	}

	grp := errgroup.Group{}
	grp.Go(func() error {
		return runGRPCServer(serverStorage, appConfig)
	})
	grp.Go(func() error {
		return runHTTPServer(serverStorage, appConfig)
	})

	return grp.Wait()
}

func runHTTPServer(serverStorage handlers.Repository, appConfig *config.Options) error {
	s := handlers.NewServer(serverStorage, appConfig)
	r := router.NewRouter(s)

	logger.Log.Infof("Build version: %v", buildVersion)
	logger.Log.Infof("Build date: %v", buildDate)
	logger.Log.Infof("Build commit: %v", buildCommit)
	logger.Log.Infow("Running server", "address", s.Config.Host)

	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	defer cancelCtx()

	srv := http.Server{Addr: s.Config.Host, Handler: r}

	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
	}()

	wg.Add(1)
	go func() {
		defer logger.Log.Info("server has been shutdown")
		defer wg.Done()
		<-ctx.Done()

		logger.Log.Info("app got a signal")
		shutdownTimeoutCtx, cancelShutdownTimeoutCtx := context.WithTimeout(context.Background(), timeoutServerShutdown)
		defer cancelShutdownTimeoutCtx()
		if err := srv.Shutdown(shutdownTimeoutCtx); err != nil {
			logger.Log.Errorf("an error occurred during server shutdown: %v", err)
		}
	}()

	var finalErr error
	if appConfig.EnableHTTPS {
		certFile, keyFile, err := mycrypto.LoadCryptoFiles()
		if err != nil {
			return err
		}
		finalErr = srv.ListenAndServeTLS(certFile, keyFile)
	} else {
		finalErr = srv.ListenAndServe()
	}

	if errors.Is(finalErr, http.ErrServerClosed) {
		return nil
	}

	return finalErr
}

func runGRPCServer(serverStorage handlers.Repository, appConfig *config.Options) error {
	serverGRPC := grpchandlers.NewServer(serverStorage, appConfig)

	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		return err
	}
	gRPCServer := grpc.NewServer()
	pb.RegisterShortenerProtoServer(gRPCServer, serverGRPC)

	fmt.Println("gRPC server is running")
	return gRPCServer.Serve(listen)
}

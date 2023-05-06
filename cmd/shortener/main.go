package main

import (
	"flag"
	"github.com/arseniy96/url-shortener/cmd/storage"
	"github.com/go-chi/chi/v5"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var serverFlags = flag.NewFlagSet("server", flag.ExitOnError)
var host *string
var resolveHost *string

func init() {
	rand.NewSource(time.Now().UnixNano())
	host = serverFlags.String("a", "localhost:8080", "server host with port")
	resolveHost = serverFlags.String("b", "http://localhost:8080", "resolve link address")
}

func main() {
	switch os.Args[1] {
	case "server":
		serverFlags.Parse(os.Args[2:])
	default:
		// PrintDefaults выводит параметры командной строки
		flag.PrintDefaults()
		os.Exit(1)
	}

	if serverFlags.Parsed() {
		if err := run(); err != nil {
			panic(err)
		}
	}
}

func run() error {
	serverStorage := storage.NewStorage()
	server := NewServer(serverStorage)

	router := chi.NewRouter()
	router.Post("/", server.CreateLink)
	router.Get("/{url_id}", server.ResolveLink)

	return http.ListenAndServe(server.config.Host, router)
}

package main

import (
	"github.com/go-chi/chi/v5"
	"math/rand"
	"net/http"
	"time"
)

const Host = "http://localhost:8080/"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	storage := NewStorage()
	server := NewServer(storage)

	router := chi.NewRouter()
	router.Post("/", server.CreateLink)
	router.Get("/{url_id}", server.ResolveLink)

	return http.ListenAndServe(`:8080`, router)
}

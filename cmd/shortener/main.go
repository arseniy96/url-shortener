package main

import (
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
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, MainHandler)
	return http.ListenAndServe(`:8080`, mux)
}

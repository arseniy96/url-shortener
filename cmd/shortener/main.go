package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const Host = "http://localhost:8080/"

var Urls map[string]string = make(map[string]string)
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

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
	mux.HandleFunc(`/`, mainHandler)
	return http.ListenAndServe(`:8080`, mux)
}

func mainHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		postHandler(writer, request)
	} else if request.Method == http.MethodGet {
		getHandler(writer, request)
	} else {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}
}

func postHandler(writer http.ResponseWriter, request *http.Request) {
	// check request
	// parse body
	// generate key
	// save url
	// response

	if request.URL.Path != "/" {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	key := generateKey()
	Urls[key] = string(body)
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("%s%s", Host, key)))
}

func getHandler(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[1:]
	url := Urls[id]
	if url == "" {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Location", url)
	writer.WriteHeader(http.StatusTemporaryRedirect)
}

func generateKey() string {
	symbols := make([]rune, 6)
	for i := range symbols {
		symbols[i] = letters[rand.Intn(len(letters))]
	}
	key := string(symbols)
	if Urls[key] != "" {
		return generateKey()
	}
	return key
}

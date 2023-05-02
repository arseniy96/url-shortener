package main

import (
	"fmt"
	"io"
	"net/http"
)

func MainHandler(storage Repository) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodPost {
			PostHandler(writer, request, storage)
		} else if request.Method == http.MethodGet {
			GetHandler(writer, request, storage)
		} else {
			http.Error(writer, "Invalid request", http.StatusBadRequest)
			return
		}
	}
}

func PostHandler(writer http.ResponseWriter, request *http.Request, storage Repository) {
	// check request
	// parse body
	// generate key
	// save url
	// response

	if request.URL.Path != "/" {
		http.Error(writer, "Invalid path", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil || len(body) == 0 {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	key := GenerateKey(storage)
	storage.Add(key, string(body))
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("%s%s", Host, key)))
}

func GetHandler(writer http.ResponseWriter, request *http.Request, storage Repository) {
	id := request.URL.Path[1:]
	url, ok := storage.Get(id)
	if !ok {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Location", url)
	writer.WriteHeader(http.StatusTemporaryRedirect)
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

var Urls = make(map[string]string)

func MainHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		PostHandler(writer, request)
	} else if request.Method == http.MethodGet {
		GetHandler(writer, request)
	} else {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}
}

func PostHandler(writer http.ResponseWriter, request *http.Request) {
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

	key := GenerateKey()
	Urls[key] = string(body)
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(fmt.Sprintf("%s%s", Host, key)))
}

func GetHandler(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[1:]
	url := Urls[id]
	if url == "" {
		http.Error(writer, "Invalid request", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Location", url)
	writer.WriteHeader(http.StatusTemporaryRedirect)
}

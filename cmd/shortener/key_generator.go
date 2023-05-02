package main

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateKey(storage Repository) string {
	symbols := make([]rune, 6)
	for i := range symbols {
		symbols[i] = letters[rand.Intn(len(letters))]
	}
	key := string(symbols)
	if _, exists := storage.Get(key); exists {
		return GenerateKey(storage)
	}
	return key
}

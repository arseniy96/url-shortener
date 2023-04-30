package main

import (
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateKey() string {
	symbols := make([]rune, 6)
	for i := range symbols {
		symbols[i] = letters[rand.Intn(len(letters))]
	}
	key := string(symbols)
	if Urls[key] != "" {
		return GenerateKey()
	}
	return key
}

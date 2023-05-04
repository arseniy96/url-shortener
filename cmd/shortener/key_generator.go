package main

import (
	"math/rand"
)

func NewGenerator(runes []rune, storage Repository) Generator {
	return Generator{
		letters: runes,
		storage: storage,
	}
}

type Generate interface {
	CreateKey() string
}

type Generator struct {
	letters []rune
	storage Repository
}

func (g Generator) CreateKey() string {
	symbols := make([]rune, 6)
	for i := range symbols {
		symbols[i] = g.letters[rand.Intn(len(g.letters))]
	}
	key := string(symbols)
	if _, exists := g.storage.Get(key); exists {
		return g.CreateKey()
	}
	return key
}

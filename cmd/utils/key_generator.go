package utils

import (
	"github.com/arseniy96/url-shortener/cmd/storage"
	"math/rand"
)

func NewGenerator(runes []rune, store storage.Repository) Generator {
	return Generator{
		letters: runes,
		storage: store,
	}
}

type Generate interface {
	CreateKey() string
}

type Generator struct {
	letters []rune
	storage storage.Repository
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

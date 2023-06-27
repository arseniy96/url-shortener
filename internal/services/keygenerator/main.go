package keygenerator

import (
	"math/rand"
)

type Repository interface {
	Get(string) (string, error)
}

func NewGenerator(store Repository) Generator {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	return Generator{
		letters: letters,
		storage: store,
	}
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
	if _, err := g.storage.Get(key); err == nil {
		return g.CreateKey()
	}
	return key
}

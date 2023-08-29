// Package keygenerator отвечает за генерацию последовательностей случайных символов.
package keygenerator

import (
	"math/rand"
)

type Repository interface {
	Get(string) (string, error)
}

// NewGenerator – функция инициализации генератора случайных строк.
func NewGenerator(store Repository) Generator {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	return Generator{
		letters: letters,
		storage: store,
	}
}

type Generator struct {
	storage Repository
	letters []rune
}

// CreateKey – функция, которая генерирует строку из случайных символов.
func (g Generator) CreateKey() string {
	size := 6
	symbols := make([]rune, size)
	for i := range symbols {
		symbols[i] = g.letters[rand.Intn(len(g.letters))]
	}
	key := string(symbols)
	if _, err := g.storage.Get(key); err == nil {
		return g.CreateKey()
	}
	return key
}

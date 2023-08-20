package keygenerator

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerator_CreateKey(t *testing.T) {
	type fields struct {
		letters []rune
		storage Repository
	}
	tests := []struct {
		name       string
		fields     fields
		wantRegexp string
	}{
		{
			name:       "should return valid key",
			wantRegexp: `^[a-zA-Z]*$`,
			fields: fields{
				letters: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
				storage: NewTestStorage(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Generator{
				letters: tt.fields.letters,
				storage: tt.fields.storage,
			}
			assert.Regexp(t, regexp.MustCompile(tt.wantRegexp), g.CreateKey(), "CreateKey()")
		})
	}
}

type TestStorage struct {
	Urls map[string]string
}

func NewTestStorage() *TestStorage {
	return &TestStorage{
		Urls: make(map[string]string),
	}
}

func (s *TestStorage) Add(_, _ string) {
	s.Urls["test"] = "Test"
}

func (s *TestStorage) Get(key string) (string, error) {
	if key == "test" {
		return "test", nil
	} else {
		return "", fmt.Errorf("Error")
	}
}

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

func TestNewGenerator(t *testing.T) {
	s := NewTestStorage()

	type args struct {
		store Repository
	}
	tests := []struct {
		name string
		args args
		want Generator
	}{
		{
			name: "success init",
			args: args{
				store: s,
			},
			want: Generator{
				letters: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
				storage: s,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewGenerator(tt.args.store), "NewGenerator(%v)", tt.args.store)
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

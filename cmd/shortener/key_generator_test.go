package main

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	tests := []struct {
		name       string
		wantRegexp string
	}{
		{
			name:       "should return valid key",
			wantRegexp: `^[a-zA-Z]*$`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := NewTestStorage()

			assert.Regexp(t, regexp.MustCompile(tt.wantRegexp), GenerateKey(storage))
		})
	}
}

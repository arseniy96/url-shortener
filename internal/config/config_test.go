package config

import (
	"reflect"
	"testing"
)

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name string
		want *Options
	}{
		{
			name: "success init",
			want: &Options{
				Host:         "localhost:8080",
				ResolveHost:  "http://localhost:8080",
				LoggingLevel: "info",
				Filename:     "/tmp/short-url-db.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitConfig(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

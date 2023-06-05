package storage

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestStorage_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// создаём объект-заглушку
	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().FindRecord(gomock.Any(), "testS").Return(Record{OriginalURL: "http://test.ru"}, nil)
	m.EXPECT().HealthCheck().Return(nil)

	type fields struct {
		Links      map[string]string
		filename   string
		dataWriter *DataWriter
		database   DatabaseInterface
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		{
			name: "success Get from database",
			fields: fields{
				database: m,
			},
			args:  args{key: "testS"},
			want:  "http://test.ru",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Links:      tt.fields.Links,
				filename:   tt.fields.filename,
				dataWriter: tt.fields.dataWriter,
				database:   tt.fields.database,
			}
			got, got1 := s.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

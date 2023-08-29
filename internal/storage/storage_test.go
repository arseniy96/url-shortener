package storage

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func TestStorage_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().FindRecord(gomock.Any(), "testS").Return(Record{OriginalURL: "http://test.ru"}, nil)
	m.EXPECT().FindRecord(gomock.Any(), "testD").Return(Record{OriginalURL: "http://test.com", DeletedFlag: true}, nil)
	links := make(map[string]string)
	links["test1"] = "http://ya.ru"

	type fields struct {
		Links    map[string]string
		database DatabaseInterface
		mode     int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success Get from database",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args:    args{key: "testS"},
			want:    "http://test.ru",
			wantErr: false,
		},
		{
			name: "success Get from database(deleted key)",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args:    args{key: "testD"},
			want:    "",
			wantErr: true,
		},
		{
			name: "success Get from memory",
			fields: fields{
				mode:  MemoryMode,
				Links: links,
			},
			args:    args{key: "test1"},
			want:    "http://ya.ru",
			wantErr: false,
		},
		{
			name: "failed Get from memory",
			fields: fields{
				mode:  MemoryMode,
				Links: links,
			},
			args:    args{key: "test2"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Links:    tt.fields.Links,
				database: tt.fields.database,
				mode:     tt.fields.mode,
			}
			got, err := s.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Restore() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStorage(t *testing.T) {
	type args struct {
		filename       string
		connectionData string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success create memory storage",
			args: args{
				filename:       "",
				connectionData: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewStorage(tt.args.filename, tt.args.connectionData)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStorage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestStorage_Restore(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().CreateDatabase().Return(nil)

	type fields struct {
		Links    map[string]string
		database DatabaseInterface
		mode     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success db restore",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			wantErr: false,
		},
		{
			name: "memory mode",
			fields: fields{
				mode: MemoryMode,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				database: tt.fields.database,
				mode:     tt.fields.mode,
			}
			if err := s.Restore(); (err != nil) != tt.wantErr {
				t.Errorf("Restore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Add(t *testing.T) {
	links := make(map[string]string)
	successUser := User{UserID: 1}
	failedUser := User{UserID: 2}
	pgErr := &pgconn.PgError{Code: pgerrcode.ForeignKeyViolation}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().FindUserByCookie(gomock.Any(), "success_user").Return(&successUser, nil)
	m.EXPECT().FindUserByCookie(gomock.Any(), "failed_user").Return(&failedUser, nil)
	m.EXPECT().SaveRecord(gomock.Any(), gomock.Any(), successUser.UserID).Return(nil)
	m.EXPECT().SaveRecord(gomock.Any(), gomock.Any(), failedUser.UserID).Return(fmt.Errorf("%w", pgErr))

	type fields struct {
		Links      map[string]string
		filename   string
		dataWriter *dataWriter
		database   DatabaseInterface
		mode       int
	}
	type args struct {
		key    string
		value  string
		cookie string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success DB mode",
			fields: fields{
				database: m,
				mode:     DBMode,
				Links:    links,
			},
			args: args{
				key:    "test",
				value:  "http://ya.ru",
				cookie: "success_user",
			},
			wantErr: false,
		},
		{
			name: "record already exists",
			fields: fields{
				database: m,
				mode:     DBMode,
				Links:    links,
			},
			args: args{
				key:    "existed",
				value:  "http://ya.ru",
				cookie: "failed_user",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Links:      tt.fields.Links,
				filename:   tt.fields.filename,
				dataWriter: tt.fields.dataWriter,
				database:   tt.fields.database,
				mode:       tt.fields.mode,
			}
			linksCount := len(s.Links)
			if err := s.Add(tt.args.key, tt.args.value, tt.args.cookie); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
			// проверяем, положилось ли в память
			if !tt.wantErr && len(s.Links) <= linksCount {
				t.Errorf("Add() failed")
			}
		})
	}
}

func TestStorage_GetByOriginURL(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().FindRecordByOriginURL(gomock.Any(), "success_url").Return(Record{ShortULR: "success_short"}, nil)
	m.EXPECT().FindRecordByOriginURL(gomock.Any(), "failed_url").Return(Record{ShortULR: ""}, fmt.Errorf("not found"))

	type fields struct {
		Links      map[string]string
		filename   string
		dataWriter *dataWriter
		database   DatabaseInterface
		mode       int
	}
	type args struct {
		originURL string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "not DB mode",
			fields: fields{
				mode: MemoryMode,
			},
			args:    args{originURL: "test"},
			want:    "",
			wantErr: true,
		},
		{
			name: "success DB mode",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args:    args{originURL: "success_url"},
			want:    "success_short",
			wantErr: false,
		},
		{
			name: "failed DB mode",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args:    args{originURL: "failed_url"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Links:      tt.fields.Links,
				filename:   tt.fields.filename,
				dataWriter: tt.fields.dataWriter,
				database:   tt.fields.database,
				mode:       tt.fields.mode,
			}
			got, err := s.GetByOriginURL(tt.args.originURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByOriginURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetByOriginURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().CreateUser(gomock.Any()).Return(&User{}, nil)

	type fields struct {
		Links      map[string]string
		filename   string
		dataWriter *dataWriter
		database   DatabaseInterface
		mode       int
	}

	tests := []struct {
		name    string
		fields  fields
		want    *User
		wantErr bool
	}{
		{
			name: "not DB mode",
			fields: fields{
				mode: MemoryMode,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success DB mode",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			want:    &User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Links:      tt.fields.Links,
				filename:   tt.fields.filename,
				dataWriter: tt.fields.dataWriter,
				database:   tt.fields.database,
				mode:       tt.fields.mode,
			}
			got, err := s.CreateUser(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().UpdateUser(gomock.Any(), 1, gomock.Any()).Return(nil)
	m.EXPECT().UpdateUser(gomock.Any(), 2, gomock.Any()).Return(fmt.Errorf("update user error"))

	type fields struct {
		Links      map[string]string
		filename   string
		dataWriter *dataWriter
		database   DatabaseInterface
		mode       int
	}
	type args struct {
		id     int
		cookie string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "not DB mode",
			fields: fields{
				mode: MemoryMode,
			},
			args: args{
				id:     0,
				cookie: "",
			},
			wantErr: true,
		},
		{
			name: "success update",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				id:     1,
				cookie: "",
			},
			wantErr: false,
		},
		{
			name: "failed update",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				id:     2,
				cookie: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Links:      tt.fields.Links,
				filename:   tt.fields.filename,
				dataWriter: tt.fields.dataWriter,
				database:   tt.fields.database,
				mode:       tt.fields.mode,
			}
			if err := s.UpdateUser(context.Background(), tt.args.id, tt.args.cookie); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_FindUserByID(t *testing.T) {
	successUser := User{UserID: 1}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().FindUserByID(gomock.Any(), 1).Return(&successUser, nil)
	m.EXPECT().FindUserByID(gomock.Any(), 2).Return(nil, fmt.Errorf("find user error"))

	type fields struct {
		Links    map[string]string
		database DatabaseInterface
		mode     int
	}
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "not DB mode",
			fields: fields{
				mode: MemoryMode,
			},
			args: args{
				userID: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success find",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				userID: 1,
			},
			want:    &successUser,
			wantErr: false,
		},
		{
			name: "unknown user",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				userID: 2,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Links:    tt.fields.Links,
				database: tt.fields.database,
				mode:     tt.fields.mode,
			}
			got, err := s.FindUserByID(context.Background(), tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetMode(t *testing.T) {
	type fields struct {
		mode int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "get mode",
			fields: fields{
				mode: DBMode,
			},
			want: DBMode,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				mode: tt.fields.mode,
			}
			if got := s.GetMode(); got != tt.want {
				t.Errorf("GetMode() = %v, want %v", got, tt.want)
			}
		})
	}
}

//nolint:dupl // it's ok
func TestStorage_HealthCheck(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().HealthCheck().Return(nil)

	type fields struct {
		database DatabaseInterface
		mode     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				database: tt.fields.database,
				mode:     tt.fields.mode,
			}
			if err := s.HealthCheck(); (err != nil) != tt.wantErr {
				t.Errorf("HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

//nolint:dupl // it's ok
func TestStorage_CloseConnection(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().Close().Return(nil)

	type fields struct {
		database DatabaseInterface
		mode     int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				database: tt.fields.database,
				mode:     tt.fields.mode,
			}
			if err := s.CloseConnection(); (err != nil) != tt.wantErr {
				t.Errorf("HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeleteUserURLs(t *testing.T) {
	successUser := User{UserID: 1}
	successRecords := []Record{{UserID: successUser.UserID}}
	successURLs := []string{"test"}
	failedURLs := []string{"test3"}
	emptyRecordsURLs := []string{"test4"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().FindUserByCookie(gomock.Any(), "success_user").Return(&successUser, nil).AnyTimes()
	m.EXPECT().FindUserByCookie(gomock.Any(), "failed_user").Return(nil, fmt.Errorf("unkoniwn user"))
	m.EXPECT().FindRecordsBatchByShortURL(gomock.Any(), successURLs).Return(successRecords, nil)
	m.EXPECT().FindRecordsBatchByShortURL(gomock.Any(), failedURLs).Return(nil, fmt.Errorf("error"))
	m.EXPECT().FindRecordsBatchByShortURL(gomock.Any(), emptyRecordsURLs).Return([]Record{}, nil)
	m.EXPECT().DeleteBatchRecords(gomock.Any(), successRecords).Return(nil)

	type fields struct {
		database DatabaseInterface
		mode     int
	}
	type args struct {
		message DeleteURLMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success delete",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				DeleteURLMessage{
					UserCookie: "success_user",
					ShortURLs:  successURLs,
				},
			},
			wantErr: false,
		},
		{
			name: "unknown user",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				DeleteURLMessage{
					UserCookie: "failed_user",
					ShortURLs:  successURLs,
				},
			},
			wantErr: true,
		},
		{
			name: "failed delete",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				DeleteURLMessage{
					UserCookie: "success_user",
					ShortURLs:  failedURLs,
				},
			},
			wantErr: true,
		},
		{
			name: "empty urls",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				DeleteURLMessage{
					UserCookie: "success_user",
					ShortURLs:  emptyRecordsURLs,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				database: tt.fields.database,
				mode:     tt.fields.mode,
			}
			if err := s.DeleteUserURLs(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserURLs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_AddBatch(t *testing.T) {
	successRecords := []Record{{ShortULR: "test", OriginalURL: "https://test.ru"}}
	failedRecords := []Record{{ShortULR: "test2", OriginalURL: "https://test2.ru"}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().SaveRecordsBatch(gomock.Any(), successRecords).Return(nil).AnyTimes()
	m.EXPECT().SaveRecordsBatch(gomock.Any(), failedRecords).Return(fmt.Errorf("error")).AnyTimes()

	type fields struct {
		Links    map[string]string
		database DatabaseInterface
		mode     int
	}
	type args struct {
		records []Record
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success in DB mode",
			fields: fields{
				Links:    make(map[string]string),
				database: m,
				mode:     DBMode,
			},
			args: args{
				records: successRecords,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Links:    tt.fields.Links,
				database: tt.fields.database,
				mode:     tt.fields.mode,
			}
			if err := s.AddBatch(context.Background(), tt.args.records); (err != nil) != tt.wantErr {
				t.Errorf("AddBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_GetByUser(t *testing.T) {
	successUser := User{UserID: 1}
	user := User{UserID: 2}
	successRecords := []Record{{UserID: successUser.UserID}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDatabaseInterface(ctrl)
	m.EXPECT().FindUserByCookie(gomock.Any(), "success_user").Return(&successUser, nil).AnyTimes()
	m.EXPECT().FindUserByCookie(gomock.Any(), "success_user2").Return(&user, nil)
	m.EXPECT().FindUserByCookie(gomock.Any(), "failed_user").Return(nil, fmt.Errorf("unkoniwn user"))
	m.EXPECT().FindRecordsByUserID(gomock.Any(), successUser.UserID).Return(successRecords, nil).AnyTimes()
	m.EXPECT().FindRecordsByUserID(gomock.Any(), user.UserID).Return(nil, fmt.Errorf("test error")).AnyTimes()

	type fields struct {
		database DatabaseInterface
		mode     int
	}
	type args struct {
		cookie string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Record
		wantErr bool
	}{
		{
			name: "not DB mode",
			fields: fields{
				database: nil,
				mode:     MemoryMode,
			},
			args: args{
				cookie: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				cookie: "success_user",
			},
			want:    successRecords,
			wantErr: false,
		},
		{
			name: "unknown user",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				cookie: "failed_user",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed get",
			fields: fields{
				database: m,
				mode:     DBMode,
			},
			args: args{
				cookie: "success_user2",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				database: tt.fields.database,
				mode:     tt.fields.mode,
			}
			got, err := s.GetByUser(context.Background(), tt.args.cookie)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

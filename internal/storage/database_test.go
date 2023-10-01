//go:build integration
// +build integration

package storage

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	"github.com/arseniy96/url-shortener/internal/logger"
)

var testDB *sql.DB
var testDatabase *Database

func TestMain(m *testing.M) {
	if err := logger.Initialize("info"); err != nil {
		panic(err)
	}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		logger.Log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		logger.Log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		logger.Log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseURL := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)

	logger.Log.Info("Connecting to database on url: ", databaseURL)

	err = resource.Expire(60) // Tell docker to hard kill the container in 120 seconds
	if err != nil {
		logger.Log.Fatalf("Could not purge resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 60 * time.Second
	if err = pool.Retry(func() error {
		testDB, err = sql.Open("postgres", databaseURL)
		if err != nil {
			return err
		}
		return testDB.Ping()
	}); err != nil {
		logger.Log.Fatalf("Could not connect to docker: %s", err)
	}

	testDatabase = &Database{DB: testDB}
	err = testDatabase.CreateDatabase()
	if err != nil {
		logger.Log.Fatalf("run migrations error: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		logger.Log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

//nolint:dupl // it's test
func TestDatabase_FindRecord(t *testing.T) {
	err := testDatabase.SaveRecord(context.Background(),
		&Record{
			UUID:        "some_id",
			ShortULR:    "test1",
			OriginalURL: "https://yan.ru",
			DeletedFlag: false,
		},
		0)
	if err != nil {
		t.Errorf("SaveRecord() error = %v", err)
		return
	}

	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    Record
		wantErr bool
	}{
		{
			name: "success find",
			args: args{
				value: "test1",
			},
			want: Record{
				UUID:        "some_id",
				ShortULR:    "test1",
				OriginalURL: "https://yan.ru",
				DeletedFlag: false,
				UserID:      0,
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				value: "not_found",
			},
			want:    Record{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testDatabase.FindRecord(context.Background(), tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//nolint:dupl // it's test
func TestDatabase_FindRecordByOriginURL(t *testing.T) {
	err := testDatabase.SaveRecord(context.Background(),
		&Record{
			UUID:        "some_id",
			ShortULR:    "test",
			OriginalURL: "https://ya.ru",
			DeletedFlag: false,
		},
		0)
	if err != nil {
		t.Errorf("SaveRecord() error = %v", err)
		return
	}

	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    Record
		wantErr bool
	}{
		{
			name: "success find",
			args: args{
				value: "https://ya.ru",
			},
			want: Record{
				UUID:        "some_id",
				ShortULR:    "test",
				OriginalURL: "https://ya.ru",
				DeletedFlag: false,
				UserID:      0,
			},
			wantErr: false,
		},
		{
			name: "not found",
			args: args{
				value: "not_found",
			},
			want:    Record{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testDatabase.FindRecordByOriginURL(context.Background(), tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_SaveRecord(t *testing.T) {
	type args struct {
		rec    *Record
		userID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "save successfully",
			args: args{
				rec: &Record{
					UUID:        "uuid",
					ShortULR:    "test",
					OriginalURL: "https://yandex.ru",
					DeletedFlag: false,
					UserID:      0,
				},
				userID: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testDatabase.SaveRecord(context.Background(), tt.args.rec, tt.args.userID); (err != nil) != tt.wantErr {
				t.Errorf("SaveRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_SaveRecordsBatch(t *testing.T) {
	type args struct {
		records []Record
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "save successfully",
			args: args{
				records: []Record{
					{
						UUID:        "uuid",
						ShortULR:    "test",
						OriginalURL: "https://google.com",
						DeletedFlag: false,
						UserID:      0,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "dublicate error",
			args: args{
				records: []Record{
					{
						UUID:        "uuid",
						ShortULR:    "test",
						OriginalURL: "https://yandex.ru",
						DeletedFlag: false,
						UserID:      0,
					},
					{
						UUID:        "uuid",
						ShortULR:    "test",
						OriginalURL: "https://yandex.ru",
						DeletedFlag: false,
						UserID:      0,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testDatabase.SaveRecordsBatch(context.Background(), tt.args.records); (err != nil) != tt.wantErr {
				t.Errorf("SaveRecordsBatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_HealthCheck(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testDatabase.HealthCheck(); (err != nil) != tt.wantErr {
				t.Errorf("HealthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_FindRecordsByUserID(t *testing.T) {
	err := testDatabase.SaveRecord(context.Background(),
		&Record{
			UUID:        "some_id_12",
			ShortULR:    "short_url_test1",
			OriginalURL: "https://something.ru",
			DeletedFlag: false,
		},
		5)
	if err != nil {
		t.Errorf("SaveRecord() error = %v", err)
		return
	}

	type args struct {
		userID int
	}
	tests := []struct {
		name        string
		args        args
		wantRecords []Record
		wantErr     bool
	}{
		{
			name: "success find",
			args: args{
				userID: 5,
			},
			wantRecords: []Record{{
				UUID:        "some_id_12",
				ShortULR:    "short_url_test1",
				OriginalURL: "https://something.ru",
				DeletedFlag: false,
				UserID:      5,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := testDatabase.FindRecordsByUserID(context.Background(), tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRecordsByUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("FindRecordsByUserID() gotRecords = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}

func TestDatabase_FindUserByCookie(t *testing.T) {
	var testCookie = "test_cookie"

	type args struct {
		cookie string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "success find",
			args: args{
				cookie: testCookie,
			},
			want: &User{
				UserID: 1,
				Cookie: testCookie,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := testDatabase.CreateUser(context.Background())
			if err != nil {
				t.Errorf("CreateUser() error = %v", err)
				return
			}
			err = testDatabase.UpdateUser(context.Background(), user.UserID, testCookie)
			if err != nil {
				t.Errorf("UpdateUser() error = %v", err)
				return
			}

			got, err := testDatabase.FindUserByCookie(context.Background(), tt.args.cookie)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindUserByCookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUserByCookie() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabase_FindUserByID(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "success find",
			args: args{
				userID: 1,
			},
			want: &User{
				Cookie: "test_cookie",
				UserID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testDatabase.FindUserByID(context.Background(), tt.args.userID)
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

func TestDatabase_FindRecordsBatchByShortURL(t *testing.T) {
	type args struct {
		urls []string
	}
	tests := []struct {
		name        string
		args        args
		wantRecords []Record
		wantErr     bool
	}{
		{
			name: "success find",
			args: args{
				urls: []string{"test198"},
			},
			wantRecords: []Record{
				{
					UUID:        "some_id",
					ShortULR:    "test198",
					OriginalURL: "https://yaasdsad.ru",
					DeletedFlag: false,
					UserID:      0,
				},
			},
			wantErr: false,
		},
	}
	//nolint:dupl // it's test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testDatabase.SaveRecord(context.Background(),
				&Record{
					UUID:        "some_id",
					ShortULR:    "test198",
					OriginalURL: "https://yaasdsad.ru",
					DeletedFlag: false,
				},
				0)
			if err != nil {
				t.Errorf("SaveRecord() error = %v", err)
				return
			}

			gotRecords, err := testDatabase.FindRecordsBatchByShortURL(context.Background(), tt.args.urls)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRecordsBatchByShortURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("FindRecordsBatchByShortURL() gotRecords = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}

func TestDatabase_DeleteBatchRecords(t *testing.T) {
	type args struct {
		records []Record
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete successfully",
			args: args{
				records: []Record{
					{
						UUID:        "some_id",
						ShortULR:    "test198",
						OriginalURL: "https://yaasdsad.ru",
						DeletedFlag: false,
						UserID:      0,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testDatabase.DeleteBatchRecords(context.Background(), tt.args.records); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBatchRecords() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabase_Close(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := testDatabase.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

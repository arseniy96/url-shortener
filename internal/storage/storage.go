package storage

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"

	"github.com/arseniy96/url-shortener/internal/logger"
)

const (
	MemoryMode = iota
	FileMode
	DBMode
)

var ErrConflict = errors.New(`already exists`)

type DatabaseInterface interface {
	FindRecord(ctx context.Context, value string) (Record, error)
	FindRecordByOriginURL(ctx context.Context, originURL string) (Record, error)
	HealthCheck() error
	Close() error
	CreateDatabase() error
	SaveRecord(context.Context, *Record, int) error
	SaveRecordsBatch(context.Context, []Record) error
	FindRecordsByUserID(context.Context, int) ([]Record, error)
	FindUserByCookie(context.Context, string) (*User, error)
	CreateUser(context.Context) (*User, error)
	UpdateUser(context.Context, int, string) error
	FindUserByID(context.Context, int) (*User, error)
}

type Storage struct {
	Links      map[string]string
	filename   string
	dataWriter *DataWriter
	database   DatabaseInterface
	mode       int
}

type Record struct {
	UUID        string `json:"uuid"`
	ShortULR    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type User struct {
	UserID int    `json:"user_id"`
	Cookie string `json:"cookie"`
}

type DataWriter struct {
	file    *os.File
	encoder *json.Encoder
}

func NewDataWriter(filename string) (*DataWriter, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Log.Error("open file error", zap.Error(err))
		return nil, err
	}

	return &DataWriter{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (w *DataWriter) WriteData(record *Record) error {
	return w.encoder.Encode(record)
}

func (w *DataWriter) Close() error {
	return w.file.Close()
}

func NewStorage(filename, connectionData string) (*Storage, error) {
	var dataWriter *DataWriter
	var db *Database
	var err error

	mode := MemoryMode

	if connectionData != "" {
		mode = DBMode
		db, err = NewDatabase(connectionData)
		if err != nil {
			return nil, err
		}
	} else if filename != "" {
		mode = FileMode
	}

	store := Storage{
		Links:      make(map[string]string),
		filename:   filename,
		dataWriter: dataWriter,
		database:   db,
		mode:       mode,
	}

	return &store, nil
}

func (s *Storage) Restore() error {
	switch s.mode {
	case FileMode:
		file, err := os.OpenFile(s.filename, os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer file.Close()

		fileScanner := bufio.NewScanner(file)

		for fileScanner.Scan() {
			var record Record
			line := fileScanner.Text()
			err = json.Unmarshal([]byte(line), &record)
			if err != nil {
				logger.Log.Error("data decoding problem", zap.Error(err))
				continue
			}

			s.Links[record.ShortULR] = record.OriginalURL
		}
	case DBMode:
		err := s.database.CreateDatabase()
		if err != nil {
			logger.Log.Error("database creation error", zap.Error(err))
			return err
		}
	}

	return nil
}

func (s *Storage) Add(key, value, cookie string) error {
	id := uuid.NewString()
	record := Record{
		UUID:        id,
		ShortULR:    key,
		OriginalURL: value,
	}

	if s.mode == DBMode {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		user, err := s.database.FindUserByCookie(ctx, cookie)
		if err != nil {
			return err
		}
		err = s.database.SaveRecord(ctx, &record, user.UserID)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return ErrConflict
			}

			logger.Log.Error("save data to database error", zap.Error(err))
			return err
		}
	} else if s.mode == FileMode {
		dataWriter, err := NewDataWriter(s.filename)
		if err != nil {
			logger.Log.Error("Open File error", zap.Error(err))
			return err
		}
		s.dataWriter = dataWriter
		defer s.dataWriter.Close()

		err = s.dataWriter.WriteData(&record)
		if err != nil {
			logger.Log.Error(zap.Error(err))
			return err
		}
	}
	s.Links[key] = value
	return nil
}

func (s *Storage) Get(key string) (string, bool) {
	if s.mode == DBMode {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		rec, err := s.database.FindRecord(ctx, key)
		if err != nil {
			return "", false
		}
		return rec.OriginalURL, true
	}
	value, found := s.Links[key]
	return value, found
}

func (s *Storage) GetByOriginURL(originURL string) (string, error) {
	if s.mode == DBMode {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		rec, err := s.database.FindRecordByOriginURL(ctx, originURL)
		if err != nil {
			return "", err
		}
		return rec.ShortULR, nil
	}

	return "", errors.New("not database mode")
}

func (s *Storage) AddBatch(ctx context.Context, records []Record) error {
	if s.mode == DBMode {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		err := s.database.SaveRecordsBatch(ctx, records)
		if err != nil {
			logger.Log.Error("save data to database error", zap.Error(err))
		}
	} else if s.mode == FileMode {
		dataWriter, err := NewDataWriter(s.filename)
		if err != nil {
			logger.Log.Error("Open File error", zap.Error(err))
		}
		s.dataWriter = dataWriter
		defer s.dataWriter.Close()
	}

	for _, record := range records {
		if s.mode == FileMode {
			err := s.dataWriter.WriteData(&record)
			if err != nil {
				logger.Log.Error(zap.Error(err))
			}
		}
		s.Links[record.ShortULR] = record.OriginalURL
	}

	return nil
}

func (s *Storage) GetByUser(ctx context.Context, cookie string) ([]Record, error) {
	if s.mode == DBMode {
		user, err := s.database.FindUserByCookie(ctx, cookie)
		if err != nil {
			return nil, err
		}
		records, err := s.database.FindRecordsByUserID(ctx, user.UserID)
		if err != nil {
			return nil, err
		}
		return records, nil
	}

	return nil, errors.New("not database mode")
}

func (s *Storage) FindUserByID(ctx context.Context, userID int) (*User, error) {
	if s.mode == DBMode {
		return s.database.FindUserByID(ctx, userID)
	}

	return nil, errors.New("not database mode")
}

func (s *Storage) CreateUser(ctx context.Context) (*User, error) {
	if s.mode == DBMode {
		return s.database.CreateUser(ctx)
	}

	return nil, errors.New("not database mode")
}

func (s *Storage) UpdateUser(ctx context.Context, id int, cookie string) error {
	if s.mode == DBMode {
		return s.database.UpdateUser(ctx, id, cookie)
	}

	return errors.New("not database mode")
}

func (s *Storage) HealthCheck() error {
	return s.database.HealthCheck()
}

func (s *Storage) CloseConnection() error {
	return s.database.Close()
}

func (s *Storage) GetMode() int {
	return s.mode
}

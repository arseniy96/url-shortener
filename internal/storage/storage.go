package storage

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
var ErrDeleted = errors.New(`was deleted`)

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
	FindRecordsBatchByShortURL(context.Context, []string) ([]Record, error)
	DeleteBatchRecords(context.Context, []Record) error
}

// Storage – структура, которая даёт доступ к хранилищу
type Storage struct {
	Links      map[string]string
	filename   string
	dataWriter *dataWriter
	database   DatabaseInterface
	mode       int
}

// Record – записать в БД
type Record struct {
	// UUID – идентификатор записи в формате uuid
	UUID string `json:"uuid"`
	// ShortULR – сокращённая ссылка
	ShortULR string `json:"short_url"`
	// OriginalURL – оригинальная ссылка
	OriginalURL string `json:"original_url"`
	// DeletedFlag – флаг, показывающий, что запись была перенесена в архив
	DeletedFlag bool `json:"is_deleted"`
	// UserID – идентификатор пользователя в системе
	UserID int `json:"user_id"`
}

// User – структура, которая хранит инфу пользователя
type User struct {
	// UserID – идентификатор пользователя в системе
	UserID int `json:"user_id"`
	// Cookie – cookie пользователя в текущей сессии
	Cookie string `json:"cookie"`
}

type dataWriter struct {
	file    *os.File
	encoder *json.Encoder
}

type DeleteURLMessage struct {
	UserCookie string
	ShortURLs  []string
}

func newDataWriter(filename string) (*dataWriter, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Log.Error("open file error", zap.Error(err))
		return nil, err
	}

	return &dataWriter{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (w *dataWriter) writeData(record *Record) error {
	return w.encoder.Encode(record)
}

func (w *dataWriter) close() error {
	return w.file.Close()
}

// NewStorage – функция инициализации хранилища
func NewStorage(filename, connectionData string) (*Storage, error) {
	var dataWriter *dataWriter
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

// Restore – функция восстановления хранилища из файла и создание БД, если это необходимо
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

// Add – сохранение оригинальной и сокращённой ссылки в хранилище
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
		dataWriter, err := newDataWriter(s.filename)
		if err != nil {
			logger.Log.Error("Open File error", zap.Error(err))
			return err
		}
		s.dataWriter = dataWriter
		defer s.dataWriter.close()

		err = s.dataWriter.writeData(&record)
		if err != nil {
			logger.Log.Error(zap.Error(err))
			return err
		}
	}
	s.Links[key] = value
	return nil
}

// Get – получение оригинальной ссылки по сокращённой
func (s *Storage) Get(key string) (string, error) {
	if s.mode == DBMode {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		rec, err := s.database.FindRecord(ctx, key)
		if err != nil {
			return "", fmt.Errorf("URL with key %v missing", key)
		}
		if rec.DeletedFlag {
			return "", ErrDeleted
		}
		return rec.OriginalURL, nil
	}
	value, found := s.Links[key]
	if !found {
		return "", ErrDeleted
	}
	return value, nil
}

// GetByOriginURL – получение сокращённой ссылки по оригинальной
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

// AddBatch – сохранение нескольких ссылок
func (s *Storage) AddBatch(ctx context.Context, records []Record) error {
	if s.mode == DBMode {
		err := s.database.SaveRecordsBatch(ctx, records)
		if err != nil {
			logger.Log.Error("save data to database error", zap.Error(err))
		}
	} else if s.mode == FileMode {
		dataWriter, err := newDataWriter(s.filename)
		if err != nil {
			logger.Log.Error("Open File error", zap.Error(err))
		}
		s.dataWriter = dataWriter
		defer s.dataWriter.close()
	}

	for _, record := range records {
		if s.mode == FileMode {
			err := s.dataWriter.writeData(&record)
			if err != nil {
				logger.Log.Error(zap.Error(err))
			}
		}
		s.Links[record.ShortULR] = record.OriginalURL
	}

	return nil
}

// GetByUser – получение всех ссылок авторизованного пользователя
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

// FindUserByID – поиск пользователя по userID
func (s *Storage) FindUserByID(ctx context.Context, userID int) (*User, error) {
	if s.mode == DBMode {
		return s.database.FindUserByID(ctx, userID)
	}

	return nil, errors.New("not database mode")
}

// CreateUser – создание нового пользователя
func (s *Storage) CreateUser(ctx context.Context) (*User, error) {
	if s.mode == DBMode {
		return s.database.CreateUser(ctx)
	}

	return nil, errors.New("not database mode")
}

// UpdateUser – сохранение cookie пользователя
func (s *Storage) UpdateUser(ctx context.Context, id int, cookie string) error {
	if s.mode == DBMode {
		return s.database.UpdateUser(ctx, id, cookie)
	}

	return errors.New("not database mode")
}

// DeleteUserURLs – удаление ссылок
func (s *Storage) DeleteUserURLs(message DeleteURLMessage) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err := s.database.FindUserByCookie(ctx, message.UserCookie)
	if err != nil {
		return err
	}
	userID := user.UserID

	var records, deletedRecords []Record

	records, err = s.database.FindRecordsBatchByShortURL(ctx, message.ShortURLs)
	if err != nil {
		return err
	}

	for _, rec := range records {
		if rec.UserID == userID {
			deletedRecords = append(deletedRecords, rec)
		}
	}

	if len(deletedRecords) == 0 {
		return nil
	}

	return s.database.DeleteBatchRecords(ctx, deletedRecords)
}

// HealthCheck – проверка работоспособности хранилища
func (s *Storage) HealthCheck() error {
	return s.database.HealthCheck()
}

// CloseConnection – закрытие соединения с БД
func (s *Storage) CloseConnection() error {
	return s.database.Close()
}

// GetMode – получение мода хранилища
func (s *Storage) GetMode() int {
	return s.mode
}

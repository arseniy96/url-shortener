package storage

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"os"
	"reflect"
	"time"
)

type DatabaseInterface interface {
	FindRecord(ctx context.Context, value string) (Record, error)
	HealthCheck() error
	Close() error
	Restore([]Record) error
}

type Storage struct {
	Links      map[string]string
	filename   string
	dataWriter *DataWriter
	database   DatabaseInterface
}

type Record struct {
	UUID        uuid.UUID `json:"uuid"`
	ShortULR    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
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

	if connectionData != "" {
		db, err = NewDatabase(connectionData)
		if err != nil {
			return nil, err
		}
	}

	store := Storage{
		Links:      make(map[string]string),
		filename:   filename,
		dataWriter: dataWriter,
		database:   db,
	}

	return &store, nil
}

func (s *Storage) Restore() error {
	file, err := os.OpenFile(s.filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	var records []Record

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		var record Record
		line := fileScanner.Text()
		err = json.Unmarshal([]byte(line), &record)
		if err != nil {
			logger.Log.Error("data decoding problem", zap.Error(err))
			continue
		}

		records = append(records, record)
		s.Links[record.ShortULR] = record.OriginalURL
	}

	if err := s.HealthCheck(); err == nil {
		err := s.database.Restore(records)
		if err != nil {
			logger.Log.Error("database restore problem", zap.Error(err))
		}
	}

	return nil
}

func (s *Storage) Add(key, value string) {
	id := uuid.New()
	record := Record{
		UUID:        id,
		ShortULR:    key,
		OriginalURL: value,
	}

	if s.filename != "" {
		dataWriter, err := NewDataWriter(s.filename)
		if err != nil {
			logger.Log.Error("Open File error", zap.Error(err))
		}
		s.dataWriter = dataWriter
		defer s.dataWriter.Close()

		s.dataWriter.WriteData(&record)
	}
	s.Links[key] = value
}

func (s *Storage) Get(key string) (string, bool) {
	if err := s.HealthCheck(); err != nil {
		logger.Log.Error("database connection error", zap.Error(err))
		value, found := s.Links[key]
		return value, found
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		rec, err := s.database.FindRecord(ctx, key)
		if err != nil {
			logger.Log.Error("database find error", zap.Error(err))
			return "", false
		}
		return rec.OriginalURL, true
	}
}

func (s *Storage) HealthCheck() error {
	if reflect.ValueOf(s.database).IsNil() {
		return fmt.Errorf("database is null")
	}
	return s.database.HealthCheck()
}

func (s *Storage) CloseConnection() error {
	return s.database.Close()
}

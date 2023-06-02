package storage

import (
	"bufio"
	"encoding/json"
	"github.com/arseniy96/url-shortener/internal/logger"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"os"
)

type Storage struct {
	Links      map[string]string
	filename   string
	dataWriter *DataWriter
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

func (p *DataWriter) WriteData(record *Record) error {
	return p.encoder.Encode(record)
}

func (p *DataWriter) Close() error {
	return p.file.Close()
}

func NewStorage(filename string) (*Storage, error) {
	var dataWriter *DataWriter

	store := Storage{
		Links:      make(map[string]string),
		filename:   filename,
		dataWriter: dataWriter,
	}

	return &store, nil
}

func (s *Storage) Restore() error {
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
	value, found := s.Links[key]
	return value, found
}

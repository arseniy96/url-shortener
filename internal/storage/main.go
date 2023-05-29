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
	Links    []Record
	filename string
	producer *Producer
}

type Record struct {
	UUID        uuid.UUID `json:"uuid"`
	ShortULR    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
}

type Producer struct {
	file    *os.File
	encoder *json.Encoder
}

func NewProducer(filename string) (*Producer, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Log.Error("open file error", zap.Error(err))
		return nil, err
	}

	return &Producer{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

func (p *Producer) WriteData(record *Record) error {
	return p.encoder.Encode(record)
}

func (p *Producer) Close() error {
	return p.file.Close()
}

func NewStorage(filename string) (*Storage, func() error, error) {
	var producer *Producer

	store := Storage{
		Links:    []Record{},
		filename: filename,
		producer: producer,
	}

	if filename != "" {
		err := store.Restore(filename)
		if err != nil {
			logger.Log.Error("restore data error", zap.Error(err))
			return &store, store.Close, nil
		}

		producer, err = NewProducer(filename)
		if err != nil {
			return nil, nil, err
		}
		store.producer = producer
	}

	return &store, store.Close, nil
}

func (s *Storage) Restore(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
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
		}

		s.Links = append(s.Links, record)
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
	s.Links = append(s.Links, record)

	if s.filename != "" {
		s.producer.WriteData(&record)
	}
}

func (s *Storage) Get(key string) (string, bool) {
	for _, link := range s.Links {
		if link.ShortULR == key {
			return link.OriginalURL, true
		}
	}

	return "", false
}

func (s *Storage) Close() error {
	return s.producer.Close()
}

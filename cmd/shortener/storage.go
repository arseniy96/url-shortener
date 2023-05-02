package main

type Storage struct {
	Urls map[string]string
}

type Repository interface {
	Add(string, string)
	Get(string) (string, bool)
}

func NewStorage() *Storage {
	return &Storage{
		Urls: make(map[string]string),
	}
}

func (s *Storage) Add(key, value string) {
	s.Urls[key] = value
}

func (s *Storage) Get(key string) (string, bool) {
	value := s.Urls[key]
	return value, value != ""
}

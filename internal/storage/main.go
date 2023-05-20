package storage

type Storage struct {
	Urls map[string]string
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
	value, found := s.Urls[key]
	return value, found
}

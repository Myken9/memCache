package inmemory

type Storage struct {
	st map[string]string
}

func NewStorage(st map[string]string) *Storage {
	return &Storage{st: st}
}

func (s *Storage) Get(key string) (value string, ok bool) {
	value, ok = s.st[key]
	return
}

func (s *Storage) Set(key, value string) {
	s.st[key] = value
}

func (s *Storage) Delete(key string) {
	delete(s.st, key)
}

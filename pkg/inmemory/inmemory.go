package inmemory

import "sync"

type Storage struct {
	st sync.Map
}

func (s *Storage) Get(key string) (value string, ok bool, err error) {
	val, ok := s.st.Load(key)
	if ok {
		value = val.(string)
	}
	return
}

func (s *Storage) Set(key, value string) (err error) {
	s.st.Store(key, value)
	return nil
}

func (s *Storage) Delete(key string) (err error) {
	s.st.Delete(key)
	return nil
}

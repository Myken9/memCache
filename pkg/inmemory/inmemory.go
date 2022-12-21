package inmemory

import (
	"memcach/pkg/server"
	"sync"
)

type Storage struct {
	st Client
	m  sync.RWMutex
}

func NewStorage(client Client) *Storage {
	return &Storage{st: client}
}

func (s *Storage) Get(key string) (items *server.Item, err error) {
	s.m.Lock()
	item, err := s.st.Get(key)
	s.m.Unlock()
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *Storage) Set(key, value string) (err error) {
	s.m.Lock()
	err = s.st.Set(server.Item{Key: key, Val: value})
	s.m.Unlock()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(key string) (err error) {
	s.m.Lock()
	err = s.st.Delete(key)
	s.m.Unlock()
	if err != nil {
		return err
	}
	return nil
}

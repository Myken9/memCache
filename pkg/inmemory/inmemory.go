package inmemory

import (
	"memcach/pkg/memcache"
)

type Storage struct {
	st *memcache.Client
}

func NewStorage(client *memcache.Client) *Storage {
	return &Storage{st: client}
}

func (s *Storage) Get(key string) (value string, err error) {
	item, err := s.st.Get(key)
	if err != nil {
		return "", err
	}
	return item.Val, nil
}

func (s *Storage) Set(key, value string) (err error) {
	err = s.st.Set(memcache.Item{Key: key, Val: value})
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Delete(key string) (err error) {
	err = s.st.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

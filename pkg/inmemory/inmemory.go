package inmemory

import (
	"github.com/bradfitz/gomemcache/memcache"
)

type Storage struct {
	st *memcache.Client
}

func NewStorage(client *memcache.Client) *Storage {
	return &Storage{st: client}
}

func (s *Storage) Get(key string) (value string, err error) {
	val, err := s.st.Get(key)
	if err != nil {
		return "", err
	}
	return string(val.Value), nil
}

func (s *Storage) Set(key, value string) (err error) {
	err = s.st.Set(&memcache.Item{Key: key, Value: []byte(value)})
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

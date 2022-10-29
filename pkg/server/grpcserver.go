package server

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"memcach/pkg/cache"
)

type CacheServer struct {
	memCache Storage
}

func NewCacheServer(st Storage) *CacheServer {
	return &CacheServer{memCache: st}
}

func (s *CacheServer) Get(_ context.Context, in *cache.Key) (*cache.Item, error) {

	value, ok, _ := s.memCache.Get(in.Key)
	if ok != true {
		err := errors.New("the cache has no values for the given key")
		return nil, err
	}
	a := cache.Item{Key: in.Key, Value: value}
	return &a, nil
}

func (s *CacheServer) Set(_ context.Context, in *cache.Item) (*empty.Empty, error) {
	err := s.memCache.Set(in.Key, in.Value)
	if err != nil {
		return nil, err
	}
	out := new(empty.Empty)
	return out, nil
}

func (s *CacheServer) Delete(_ context.Context, in *cache.Key) (*empty.Empty, error) {
	err := s.memCache.Delete(in.Key)
	if err != nil {
		return nil, err
	}
	out := new(empty.Empty)
	return out, nil
}

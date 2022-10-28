package server

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"memcach/pkg/cache"
)

type CacheServer struct {
	memCache map[string]string
}

func (g *CacheServer) Get(ctx context.Context, in *cache.Key) (*cache.Item, error) {

	_, ok := g.memCache[in.Key]
	if ok != true {
		err := errors.New("the cache has no values for the given key")
		return nil, err
	}
	a := cache.Item{Key: in.Key, Value: g.memCache[in.Key]}
	return &a, nil
}

func (g *CacheServer) Set(ctx context.Context, in *cache.Item) (*empty.Empty, error) {
	if g.memCache == nil {
		g.memCache = make(map[string]string)
	}
	g.memCache[in.Key] = in.Value
	out := new(empty.Empty)
	return out, nil
}

func (g *CacheServer) Delete(ctx context.Context, in *cache.Key) (*empty.Empty, error) {
	delete(g.memCache, in.Key)
	out := new(empty.Empty)
	return out, nil
}

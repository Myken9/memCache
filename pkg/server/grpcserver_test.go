package server

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"memcach/pkg/cache"
	"testing"
	"time"
)

func TestCacheServer_Set(t *testing.T) {
	s := initServer()
	t.Run("Add successful", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		item, err := s.Set(ctx, &cache.Item{Key: "123", Value: "some string"})
		assert.Nil(t, err)
		assert.Equal(t, new(empty.Empty), item)
	})
}

func TestCacheServer_Get(t *testing.T) {
	s := initServer()
	t.Run("Get successful", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		item, err := s.Get(ctx, &cache.Key{Key: "777"})
		assert.Nil(t, err)
		assert.Equal(t, &cache.Item{Key: "777", Value: "777 string"}, item)
	})

	t.Run("Get unsuccessful", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		item, err := s.Get(ctx, &cache.Key{Key: "123"})
		assert.Nil(t, item)
		assert.Equal(t, errors.New("the cache has no values for the given key"), err)
	})
}

func TestCacheServer_Delete(t *testing.T) {
	s := initServer()
	t.Run("Delete successful", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		item, err := s.Delete(ctx, &cache.Key{Key: "777"})
		assert.Nil(t, err)
		assert.Equal(t, new(empty.Empty), item)

	})

	t.Run("Delete successful", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		item, err := s.Get(ctx, &cache.Key{Key: "777"})
		assert.Nil(t, item)
		assert.Equal(t, errors.New("the cache has no values for the given key"), err)

	})
}

func initServer() *CacheServer {
	memCache := &CacheServer{memCache: make(map[string]string)}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	memCache.Set(ctx, &cache.Item{Key: "777", Value: "777 string"})
	return memCache
}

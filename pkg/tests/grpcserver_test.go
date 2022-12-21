package tests

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"memcach/pkg/cache"
	"memcach/pkg/inmemory"
	"memcach/pkg/memcache"
	"memcach/pkg/server"
	"testing"
	"time"
)

func TestCacheServer_Set(t *testing.T) {
	s := initServer()
	t.Run("Set successful", func(t *testing.T) {
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
		item, err := s.Get(ctx, &cache.Key{Key: "12322"})
		assert.Nil(t, item)
		assert.Equal(t, errors.New("memcached: cache miss"), err)
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
		assert.Equal(t, errors.New("memcached: cache miss"), err)

	})
}

func initServer() *server.CacheServer {
	mc, err := memcache.New(":11211")
	if err != nil {
		panic(err)
	}
	ns := inmemory.NewStorage(mc)
	srv := server.NewCacheServer(ns)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = srv.Set(ctx, &cache.Item{Key: "777", Value: "777 string"})
	if err != nil {
		panic(err)
	}
	return srv
}

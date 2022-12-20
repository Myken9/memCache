package memcache

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"memcach/pkg/server"
	"testing"
)

func TestClient_Set(t *testing.T) {
	client, _ := New(":11211")
	t.Run("Set successful", func(t *testing.T) {
		err := client.Set(server.Item{Key: "qwe", Val: "asdasd234234"})
		assert.Nil(t, err)
		err = client.Set(server.Item{Key: "123", Val: "123123123123dfsdfb\r\n"})
		assert.Nil(t, err)
	})
}

func TestClient_Get(t *testing.T) {
	client, _ := New(":11211")
	t.Run("Set successful", func(t *testing.T) {
		get, err := client.Get("123")
		assert.Nil(t, err)
		assert.Equal(t, server.Item{Key: "123", Val: "123123123123dfsdfb\r\n"}, get)
		get, err = client.Get("qwe")
		assert.Nil(t, err)
		assert.Equal(t, server.Item{Key: "qwe", Val: "asdasd234234"}, get)
	})
}

func TestClient_Delete(t *testing.T) {
	client, _ := New(":11211")
	t.Run("Set successful", func(t *testing.T) {
		err := client.Delete("123")
		assert.Nil(t, err)
		err = client.Delete("qwe")
		assert.Nil(t, err)

		get, err := client.Get("123")
		assert.Equal(t, errors.New("memcached: cache miss"), err)
		assert.Equal(t, server.Item{}, get)

		get, err = client.Get("qwe")
		assert.Equal(t, errors.New("memcached: cache miss"), err)
		assert.Equal(t, server.Item{}, get)
	})
}

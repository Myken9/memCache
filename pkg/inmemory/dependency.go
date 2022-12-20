package inmemory

import "memcach/pkg/server"

type Client interface {
	Close()
	Get(key string) (i server.Item, err error)
	Set(item server.Item) (err error)
	Delete(key string) error
}

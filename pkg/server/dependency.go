package server

type Storage interface {
	Set(key, value string) (err error)
	Get(key string) (item *Item, err error)
	Delete(key string) (err error)
}

type Item struct {
	Key string
	Val string
}

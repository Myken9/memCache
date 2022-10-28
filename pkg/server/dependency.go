package server

type Storage interface {
	Set(key, value string) (err error)
	Get(key string) (value string, ok bool, err error)
	Delete(key string) (err error)
}

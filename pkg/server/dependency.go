package server

type Storage interface {
	Set(key, value string)
	Get(key string) (value string, ok bool)
	Delete(key string)
}

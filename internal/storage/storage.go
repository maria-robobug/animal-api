package storage

// Cache describes a generic caching interface
type Cache interface {
	Get(key string) (interface{}, error)
}

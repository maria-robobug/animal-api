package storage

// Cache describes a generic caching interface
type Cache interface {
	Get(key string) (interface{}, error)
	// Put(key string, val interface{}, timeout time.Duration) error
	// Delete(key string) error
}
package storage

// Repository describes a generic data storage interface
type Repository interface {
	Save()
	Get()
	Delete()
}

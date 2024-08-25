package repository

type BaseRepository[T any] interface {
	Create(payload *T) error
	// List() ([]*T, error)
	// Get(id string) (*T, error)
	// Update(id string, payload *T) (*T, error)
	// Delete(id string) (*T, error)
}

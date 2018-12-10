package store

// Store ...
type Store interface {
	Save(v interface{}) error
	Edit(v interface{}, key string) error
	Delete(v interface{}) error
	GetBySomeField(v ...interface{}) (interface{}, error)
}

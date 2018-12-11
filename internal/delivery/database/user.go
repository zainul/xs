package database

import "github.com/zainul/xs/internal/entity"

// UserStore ...
type UserStore interface {
	Save(user entity.User) error
	Edit(user entity.User, key string) error
	Delete(key interface{}) error
	GetBySomeField(v ...interface{}) (interface{}, error)
}

package repository

import (
	"github.com/zainul/xs/internal/entity"
)

// UserRepository ...
type UserRepository interface {
	Save(user entity.User) error
	Edit(user entity.User, email string) error
	GetByField(field interface{}, fieldName string) *entity.User
	GenerateAccountNumber() string
}

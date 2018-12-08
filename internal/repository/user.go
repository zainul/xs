package repository

import (
	"github.com/zainul/xs/internal/domain"
)

// UserRepository ...
type UserRepository interface {
	Save(user domain.User) error
	Edit(user domain.User, email string) error
	GetByField(field interface{}, fieldName string) *domain.User
}

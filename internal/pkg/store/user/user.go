package user

import (
	"github.com/zainul/xs/internal/entity"
	"github.com/zainul/xs/internal/pkg/store"
)

// Store is represented of user process
type Store interface {
	store.Store
	SomeCustomAction(field string) (*entity.User, error)
}

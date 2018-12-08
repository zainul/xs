package postgres

import (
	"database/sql"
	"fmt"

	"github.com/zainul/xs/internal/domain"
	"github.com/zainul/xs/internal/pkg/error/dberror"
	"github.com/zainul/xs/internal/repository"
)

type user struct {
	DB *sql.DB
}

// NewUserRepository ...
func NewUserRepository(conn *sql.DB) repository.UserRepository {
	return &user{
		DB: conn,
	}
}

func (u *user) Save(user domain.User) error {
	query := `insert into users set()`
	_, err := u.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("%v %v", dberror.FailedPrepareQuery, err)
	}

	return nil
}
func (u *user) Edit(user domain.User, email string) error {
	return nil
}
func (u *user) GetByField(field interface{}, fieldName string) *domain.User {
	return nil
}

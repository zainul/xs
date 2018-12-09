package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/zainul/xs/internal/entity"
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

func (u *user) Save(user entity.User) error {
	query := `INSERT INTO users (
		email, 
		username, 
		balance, 
		first_name, 
		last_name, 
		phone_number, 
		point, 
		citizen_id, 
		refferal_code, 
		account_number, 
		created_at, 
		status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	stmt, err := u.DB.Prepare(query)

	if err != nil {
		return fmt.Errorf("%v %v", dberror.FailedPrepareQuery, err)
	}

	_, err = stmt.Exec(
		user.Email,
		user.Username,
		0,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		0,
		user.CitizenID,
		user.RefferalCode,
		user.AccountNumber,
		time.Now(),
		0,
	)

	stmt.Close()

	if err != nil {
		return fmt.Errorf("%v %v", dberror.FailedToExecQuery, err)
	}

	return nil
}
func (u *user) Edit(user entity.User, email string) error {
	return nil
}
func (u *user) GetByField(field interface{}, fieldName string) *entity.User {
	return nil
}

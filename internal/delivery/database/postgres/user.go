package postgres

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/zainul/xs/internal/entity"
	"github.com/zainul/xs/internal/pkg/error/dberror"
	"github.com/zainul/xs/internal/repository"
)

type user struct {
	DB *sql.DB
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
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

func (u *user) GenerateAccountNumber() string {
	// for temporary just random string, should be in storage
	return randomString(25)
}

// for temporary just random string, should be in storage
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

// for temporary just random string, should be in storage
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

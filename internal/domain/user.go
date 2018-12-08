package domain

import (
	"time"
)

// User is represented of the user
type User struct {
	Email         string    `json:"email"`
	Username      string    `json:"username"`
	Balance       float64   `json:"balance"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	PhoneNumber   string    `json:"phone_number"`
	Point         float64   `json:"point"`
	CitizenID     string    `json:"citizen_id"`
	RefferalCode  string    `json:"refferal_code"`
	AccountNumber string    `json:"account_number"`
	CreatedAt     time.Time `json:"created_at"`
	Status        int64     `json:"status"`
}

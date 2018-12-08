package types

// UserRegister ...
type UserRegister struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_number"`
	CitizenID    string `json:"citizen_id"`
	RefferalCode string `json:"refferal_code"`
}

package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zainul/xs/internal/usecase"
)

// Userhandler ...
type Userhandler struct {
	UserUsecase usecase.User
}

// NewUserHandler ...
func NewUserHandler(
	route *mux.Router, uUser usecase.User) {

	handler := &Userhandler{
		UserUsecase: uUser,
	}

	route.HandleFunc("/users", handler.Register).Methods(http.MethodPost)
}

// Register ...
func (u *Userhandler) Register(w http.ResponseWriter, r *http.Request) {
	// u.UserUsecase.Register()
}

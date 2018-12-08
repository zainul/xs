package http

import (
	"encoding/json"
	"net/http"

	"github.com/zainul/xs/internal/delivery/http/types"
	"github.com/zainul/xs/internal/domain"
	"github.com/zainul/xs/internal/pkg/error/deliveryerror"

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

	route.HandleFunc("/register", handler.Register).Methods(http.MethodPost)
}

// Register ...
func (u *Userhandler) Register(w http.ResponseWriter, r *http.Request) {
	response := Response{}
	body := types.UserRegister{}
	user := domain.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)

	if err != nil {
		err := deliveryerror.GetError(deliveryerror.BadRequest, err)
		response.Error = *err

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	user.CitizenID = body.CitizenID
	user.Email = body.Email
	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.PhoneNumber = body.PhoneNumber
	user.RefferalCode = body.RefferalCode
	user.Username = body.Username

	errDelivery := u.UserUsecase.Register(user)

	if errDelivery != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Error = *errDelivery
	}

	json.NewEncoder(w).Encode(response)
}

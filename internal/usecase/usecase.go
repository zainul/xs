package usecase

import (
	"github.com/zainul/xs/internal/entity"
	"github.com/zainul/xs/internal/pkg/error/deliveryerror"
)

// User user usecase is activity of user
type User interface {
	Register(user entity.User) *deliveryerror.Error
	EditProfile(user entity.User) *deliveryerror.Error
	ResetPassword(email string) *deliveryerror.Error
	DefaultInfo(accountNumber string) *entity.User
}

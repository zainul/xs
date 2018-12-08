package usecase

import (
	"github.com/zainul/xs/internal/domain"
	"github.com/zainul/xs/internal/pkg/error/deliveryerror"
)

// User user usecase is activity of user
type User interface {
	Register(user domain.User) *deliveryerror.Error
	EditProfile(user domain.User) *deliveryerror.Error
	ResetPassword(email string) *deliveryerror.Error
	DefaultInfo(accountNumber string) *domain.User
}

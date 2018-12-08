package validator

import (
	"errors"
	"fmt"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate ...
func Validate(obj interface{}) error {

	err := validate.Struct(obj)

	errStrs := make([]string, 0)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			errStrs = append(errStrs, fmt.Sprintf("%v %v", err.StructField(), err.Tag()))
		}

		return errors.New(strings.Join(errStrs, " , "))
	}

	return nil
}

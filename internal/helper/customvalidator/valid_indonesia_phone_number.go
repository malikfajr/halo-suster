package customvalidator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func validIdPhone(fl validator.FieldLevel) bool {
	// Regular expression untuk memeriksa URL HTTP(s) yang memiliki domain lengkap
	regex := regexp.MustCompile(`^\+62\d{8,13}$`)

	return regex.MatchString(fl.Field().String())
}

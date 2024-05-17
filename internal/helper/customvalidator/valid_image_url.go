package customvalidator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func validImageUrl(fl validator.FieldLevel) bool {
	// Regular expression untuk memeriksa URL HTTP(s) yang memiliki domain lengkap
	regex := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}/\S+\.(jpg|jpeg|png|gif)$`)

	return regex.MatchString(fl.Field().String())
}

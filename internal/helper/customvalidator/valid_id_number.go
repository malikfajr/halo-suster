package customvalidator

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func validIdNumber(fl validator.FieldLevel) bool {
	v := fl.Field().Int()

	IdNumber := strconv.Itoa(int(v))
	if len(IdNumber) != 16 {
		log.Println("length not 16")
		return false
	}

	return true
}

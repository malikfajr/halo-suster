package customvalidator

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewCustomValidator(validator *validator.Validate) *customValidator {
	cv := &customValidator{
		validator: validator,
	}

	cv.validator.RegisterValidation("imageUrl", validImageUrl)
	cv.validator.RegisterValidation("nip", validNip)

	return cv
}

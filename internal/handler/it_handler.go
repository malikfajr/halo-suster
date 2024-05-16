package handler

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/entity"
	"github.com/malikfajr/halo-suster/internal/exception"
	"github.com/malikfajr/halo-suster/internal/helper/jwt"
	"github.com/malikfajr/halo-suster/internal/repository"
	"github.com/malikfajr/halo-suster/internal/usecase"
)

func NewItHandler(pool *pgxpool.Pool) iITHandler {
	return &ITHandler{
		pool: pool,
	}
}

type ITHandler struct {
	pool *pgxpool.Pool
}

//
func (i *ITHandler) Login(e echo.Context) error {
	payload := &entity.ITStaffLogin{}
	if err := e.Bind(payload); err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewBadRequest("request does't pass validation"))
	}

	if err := e.Validate(payload); err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewBadRequest("request does't pass validation"))
	}

	itUsecase := usecase.NewItUsecase(i.pool, repository.NewItRepository())

	data, err := itUsecase.Login(e.Request().Context(), payload)
	if err != nil {
		ce, ok := err.(*exception.CustomError)
		if ok {
			return e.JSON(ce.StatusCode, ce)
		}
		panic(err)
	}

	return e.JSON(http.StatusOK, &jsonOk{
		Message: "User registered successfully",
		Data: &entity.ITStaffResponse{
			ID:          data.ID,
			Nip:         data.Nip,
			Name:        data.Name,
			AccessToken: jwt.CreateToken(data.ID, "it"),
		},
	})
}

//
func (i *ITHandler) Register(e echo.Context) error {
	payload := &entity.ITStaffRegister{}
	if err := e.Bind(payload); err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewBadRequest("request does't pass validation"))
	}

	if err := e.Validate(payload); err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewBadRequest("request does't pass validation"))
	}

	itUsecase := usecase.NewItUsecase(i.pool, repository.NewItRepository())

	data, err := itUsecase.Register(e.Request().Context(), payload)
	if err != nil {
		ce, ok := err.(*exception.CustomError)
		if ok {
			return e.JSON(ce.StatusCode, ce)
		}
		panic(err)
	}

	return e.JSON(http.StatusOK, &jsonOk{
		Message: "User registered successfully",
		Data: &entity.ITStaffResponse{
			ID:          data.ID,
			Nip:         data.Nip,
			Name:        data.Name,
			AccessToken: jwt.CreateToken(data.ID, "it"),
		},
	})
}

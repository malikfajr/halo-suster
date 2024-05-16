package handler

import (
	"context"
	"net/http"
	"strconv"

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
	payload := &entity.AuthLogin{}
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
		Data: &entity.AuthResponse{
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

	return e.JSON(http.StatusCreated, &jsonOk{
		Message: "User registered successfully",
		Data: &entity.AuthResponse{
			ID:          data.ID,
			Nip:         data.Nip,
			Name:        data.Name,
			AccessToken: jwt.CreateToken(data.ID, "it"),
		},
	})
}

func (i *ITHandler) GetAllUsers(e echo.Context) error {
	params := &entity.UserParam{}

	if userId := e.QueryParam("userId"); userId != "" {
		params.UserId = userId
	}

	if name := e.QueryParam("name"); name != "" {
		params.Name = name
	}

	if nip := e.QueryParam("nip"); nip != "" {
		params.Nip = nip
	}

	if role := e.QueryParam("role"); role != "" {
		if i.validRole(role) {
			params.Role = role
		}
	}

	if createdAt := e.QueryParam("createdAt"); createdAt != "" {
		if i.validOrder(createdAt) {
			params.CreatedAt = createdAt
		}
	}

	if n, err := strconv.Atoi(e.QueryParam("limit")); err != nil {
		params.Limit = 5
	} else {
		if n < 0 {
			params.Limit = 5
		} else {
			params.Limit = n

		}
	}

	if n, err := strconv.Atoi(e.QueryParam("offset")); err != nil {
		params.Offset = 0
	} else {
		if n < 0 {
			params.Offset = 0
		} else {
			params.Offset = n
		}
	}

	itUsecase := usecase.NewItUsecase(i.pool, repository.NewItRepository())

	data := itUsecase.GetAllUser(context.Background(), params)

	return e.JSON(http.StatusOK, &jsonOk{
		Message: "success",
		Data:    data,
	})

}

func (i *ITHandler) validOrder(key string) bool {
	orders := map[string]bool{
		"asc":  true,
		"desc": true,
	}

	_, ok := orders[key]
	return ok
}

func (i *ITHandler) validRole(key string) bool {
	roles := map[string]bool{
		"it":    true,
		"nurse": true,
	}

	_, ok := roles[key]
	return ok
}

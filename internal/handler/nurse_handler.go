package handler

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/entity"
	"github.com/malikfajr/halo-suster/internal/exception"
	"github.com/malikfajr/halo-suster/internal/helper/jwt"
	"github.com/malikfajr/halo-suster/internal/usecase"
)

type hNurse struct {
	pool *pgxpool.Pool
}

// Update implements nurseHanlder.
func (n *hNurse) Update(e echo.Context) error {
	payload := &entity.EditNursePayload{}

	if err := e.Bind(payload); err != nil {
		return e.JSON(400, exception.NewBadRequest("request does't pass validation"))
	}

	if err := e.Validate(payload); err != nil {
		return e.JSON(400, exception.NewBadRequest("request does't pass validation"))
	}

	nc := usecase.NewNurseCase(n.pool)

	err := nc.Update(context.Background(), payload)
	if err != nil {
		ex, ok := err.(*exception.CustomError)
		if ok {
			return e.JSON(ex.StatusCode, ex)
		}
		panic(err)
	}

	return e.JSON(200, &jsonOk{
		Message: "success",
	})
}

// AddAccess implements nurseHanlder.
func (n *hNurse) AddAccess(e echo.Context) error {
	type Payload struct {
		Password string `json:"password" validate:"required,min=5,max=33"`
	}

	payload := &Payload{}

	userId := e.Param("userId")

	if err := e.Bind(payload); err != nil {
		return e.JSON(400, exception.NewBadRequest("request doesn't pass validation"))
	}

	if err := e.Validate(payload); err != nil {
		return e.JSON(400, exception.NewBadRequest("request doesn't pass validation"))
	}

	nc := usecase.NewNurseCase(n.pool)

	nc.AddAccess(context.Background(), userId, payload.Password)

	return e.JSON(200, &jsonOk{
		Message: "success",
	})
}

// Destroy implements nurseHanlder.
func (n *hNurse) Destroy(e echo.Context) error {
	userId := e.Param("userId")

	nc := usecase.NewNurseCase(n.pool)

	err := nc.Delete(context.Background(), userId)
	if err != nil {
		ex, ok := err.(*exception.CustomError)
		if ok {
			return e.JSON(ex.StatusCode, ex)
		}
		panic(err)
	}

	return e.JSON(200, &jsonOk{
		Message: "success",
	})
}

// GetById implements nurseHanlder.
func (*hNurse) GetById(e echo.Context) error {
	panic("unimplemented")
}

// Login implements nurseHanlder.
func (n *hNurse) Login(e echo.Context) error {
	payload := &entity.AuthLogin{}

	if err := e.Bind(payload); err != nil {
		return e.JSON(400, exception.NewBadRequest("request doesn't pass validation"))
	}

	if err := e.Validate(payload); err != nil {
		return e.JSON(400, exception.NewBadRequest("request doesn't pass validation"))
	}

	nc := usecase.NewNurseCase(n.pool)

	user, err := nc.Login(context.Background(), payload)
	if err != nil {
		ex, ok := err.(*exception.CustomError)
		if ok {
			return e.JSON(ex.StatusCode, ex)
		}
		panic(err)
	}

	return e.JSON(200, &jsonOk{
		Message: "success",
		Data: &entity.AuthResponse{
			ID:          user.ID,
			Nip:         user.Nip,
			Name:        user.Name,
			AccessToken: jwt.CreateToken(user.ID, user.Nip, "nurse"),
		},
	})
}

// Register implements nurseHanlder.
func (n *hNurse) Register(e echo.Context) error {
	payload := &entity.AddNursePayload{}
	if err := e.Bind(payload); err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewBadRequest("request does't pass validation"))
	}

	if err := e.Validate(payload); err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewBadRequest("request does't pass validation"))
	}

	// nCase := usecase.NewNurseCase()
	nUsecase := usecase.NewNurseCase(n.pool)

	err := nUsecase.Register(context.Background(), payload)

	if err != nil {
		ce, ok := err.(*exception.CustomError)
		if ok {
			return e.JSON(ce.StatusCode, ce)
		}
		panic(err)
	}

	return e.JSON(http.StatusCreated, &jsonOk{
		Message: "User registered successfully",
		Data: &entity.NurseResponse{
			UserId: payload.Id,
			Name:   payload.Name,
			Nip:    payload.Nip,
		},
	})
}

func NewNurseHandler(pool *pgxpool.Pool) nurseHanlder {
	return &hNurse{
		pool: pool,
	}
}

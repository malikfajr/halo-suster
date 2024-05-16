package handler

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	echo "github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/entity"
	"github.com/malikfajr/halo-suster/internal/exception"
	"github.com/malikfajr/halo-suster/internal/usecase"
)

type hNurse struct {
	pool *pgxpool.Pool
}

// AddAccess implements nurseHanlder.
func (*hNurse) AddAccess(e echo.Context) error {
	panic("unimplemented")
}

// Destroy implements nurseHanlder.
func (*hNurse) Destroy(e echo.Context) error {
	panic("unimplemented")
}

// GetById implements nurseHanlder.
func (*hNurse) GetById(e echo.Context) error {
	panic("unimplemented")
}

// Login implements nurseHanlder.
func (*hNurse) Login(e echo.Context) error {
	panic("unimplemented")
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

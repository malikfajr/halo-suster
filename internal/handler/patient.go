package handler

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/entity"
	"github.com/malikfajr/halo-suster/internal/exception"
	"github.com/malikfajr/halo-suster/internal/repository"
	"github.com/malikfajr/halo-suster/internal/usecase"
)

type PatientHandler struct {
	pool *pgxpool.Pool
}

func (p *PatientHandler) Record(e echo.Context) error {
	payload := &entity.AddPatientPayload{}

	if err := e.Bind(payload); err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewBadRequest("payload not valid"))
	}

	if err := e.Validate(payload); err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewBadRequest("request doesn't pass validation"))
	}

	pCase := usecase.NewPatientCase(p.pool, &repository.PatientRepository{})

	patient, err := pCase.Insert(context.TODO(), payload)
	if err != nil {
		panic(err)
	}

	return e.JSON(http.StatusCreated, patient)
}

func (p *PatientHandler) GetAll(e echo.Context) error {
	params := &entity.PatientQueryParam{}
	e.Bind(params)

	pCase := usecase.NewPatientCase(p.pool, &repository.PatientRepository{})
	patients := pCase.GetAll(context.TODO(), params)

	return e.JSON(200, &jsonOk{
		Message: "success",
		Data:    patients,
	})
}

func NewPatientHandler(pool *pgxpool.Pool) *PatientHandler {
	return &PatientHandler{
		pool: pool,
	}
}

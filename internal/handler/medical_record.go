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

type MedicalRecordHanlder struct {
	Pool *pgxpool.Pool
}

func (m *MedicalRecordHanlder) Insert(e echo.Context) error {
	user, _ := e.Get("user").(*jwt.JWTClaim)

	payload := &entity.AddMedicalRecordPayload{}

	if err := e.Bind(payload); err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewBadRequest("reques doesn't pass validation"))
	}

	if err := e.Validate(payload); err != nil {
		return e.JSON(http.StatusBadRequest, exception.NewBadRequest("request doesn't pass validation"))
	}

	payload.UserId = user.ID
	payload.UserNip = user.Nip

	mrCase := usecase.NewMedicalCase(m.Pool)
	err := mrCase.Insert(context.TODO(), payload)
	if err != nil {
		ex, ok := err.(*exception.CustomError)
		if ok {
			return e.JSON(ex.StatusCode, ex)
		}
		panic(err)
	}

	return e.JSON(http.StatusCreated, &jsonOk{
		Message: "success",
	})

}

func (m *MedicalRecordHanlder) GetAll(e echo.Context) error {
	params := &entity.MedicalRecordQueryParam{}
	e.Bind(params)

	mrCase := usecase.NewMedicalCase(m.Pool)
	records := mrCase.GetAll(context.TODO(), params)

	return e.JSON(200, &jsonOk{
		Message: "success",
		Data:    records,
	})
}

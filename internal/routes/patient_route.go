package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/internal/handler"
	"github.com/malikfajr/halo-suster/internal/middleware"
)

func newPatientRoute(e *echo.Echo, pool *pgxpool.Pool) {
	h := handler.NewPatientHandler(pool)

	e.POST("/v1/medical/patient", middleware.Auth(h.Record, "it", "nurse"))
	e.GET("/v1/medical/patient", h.GetAll)
}

package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/internal/handler"
	"github.com/malikfajr/halo-suster/internal/middleware"
)

func newMedicalRecordRoute(e *echo.Echo, pool *pgxpool.Pool) {
	h := &handler.MedicalRecordHanlder{Pool: pool}

	e.POST("/v1/medical/record", middleware.Auth(h.Insert, "it", "nurse"))
	e.GET("/v1/medical/record", middleware.Auth(h.GetAll, "it", "nurse"))
}

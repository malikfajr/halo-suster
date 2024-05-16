package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/internal/handler"
	"github.com/malikfajr/halo-suster/internal/middleware"
)

func newNurseRoute(e *echo.Echo, pool *pgxpool.Pool) {

	h := handler.NewNurseHandler(pool)

	g := e.Group("/v1/user/nurse")

	g.POST("/register", middleware.Auth(h.Register, "it"))
}

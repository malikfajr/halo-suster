package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/internal/handler"
	"github.com/malikfajr/halo-suster/internal/middleware"
)

func newITRoute(e *echo.Echo, pool *pgxpool.Pool) {

	h := handler.NewItHandler(pool)

	e.GET("/v1/user", middleware.Auth(h.GetAllUsers, "it"))

	g := e.Group("/v1/user/it")

	g.POST("/register", h.Register)
	g.POST("/login", h.Login)
}

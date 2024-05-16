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
	g.PUT("/:userId", middleware.Auth(h.Update, "it"))
	g.DELETE("/:userId", middleware.Auth(h.Destroy, "it"))
	g.POST("/:userId/access", middleware.Auth(h.AddAccess, "it"))
}

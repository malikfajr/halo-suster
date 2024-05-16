package routes

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

func NewRoute(e *echo.Echo, pool *pgxpool.Pool) {
	newITRoute(e, pool)
}

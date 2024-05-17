package internal

import (
	"context"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/malikfajr/halo-suster/config"
	"github.com/malikfajr/halo-suster/internal/driver/db"
	"github.com/malikfajr/halo-suster/internal/helper/customvalidator"
	"github.com/malikfajr/halo-suster/internal/routes"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	cv := customvalidator.NewCustomValidator(validator.New())

	e.Validator = cv

	dbAddress := config.GetDBAdd()
	pool := db.NewPool(context.Background(), dbAddress)

	routes.NewRoute(e, pool)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

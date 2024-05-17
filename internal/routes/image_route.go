package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/malikfajr/halo-suster/internal/handler"
	"github.com/malikfajr/halo-suster/internal/middleware"
)

func newImageRoute(e *echo.Echo) {
	h := &handler.ImageHandler{}

	e.POST("/v1/image", middleware.Auth(h.UploadImage, "it", "nurse"))
}

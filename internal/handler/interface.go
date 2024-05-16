package handler

import "github.com/labstack/echo/v4"

type iITHandler interface {
	Login(e echo.Context) error
	Register(e echo.Context) error
}

type jsonOk struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

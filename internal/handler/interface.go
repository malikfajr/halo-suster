package handler

import "github.com/labstack/echo/v4"

type iITHandler interface {
	Login(e echo.Context) error
	Register(e echo.Context) error
	GetAllUsers(e echo.Context) error
}

type nurseHanlder interface {
	Login(e echo.Context) error
	Register(e echo.Context) error
	GetById(e echo.Context) error
	Destroy(e echo.Context) error
	AddAccess(e echo.Context) error
	Update(e echo.Context) error
}

type jsonOk struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

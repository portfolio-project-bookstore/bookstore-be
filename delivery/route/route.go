package route

import (
	_user "bookstore/delivery/controller/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc *_user.UserController) {

	// route users
	e.POST("/signup", uc.Create())
	e.GET("/users", uc.GetAll())
	e.GET("/users/:id", uc.GetById())
	e.PUT("/users/:id", uc.Update())
	e.DELETE("/users/:id", uc.Delete())
}

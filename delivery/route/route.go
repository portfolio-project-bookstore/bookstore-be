package route

import (
	_user "bookstore/delivery/controller/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, cc *_user.UserController) {

	// route users
	e.POST("/users", cc.Create())
	e.GET("/users", cc.GetAll())
	e.GET("/users/:id", cc.GetById())
	e.PUT("/users/:id", cc.Update())
	e.DELETE("/users/:id", cc.Delete())
}

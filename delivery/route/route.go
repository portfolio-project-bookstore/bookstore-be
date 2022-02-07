package route

import (
	_auth "bookstore/delivery/controller/auth"
	_user "bookstore/delivery/controller/user"
	_middleware "bookstore/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uc *_user.UserController, ac *_auth.AuthController) {
	// route login
	e.POST("/signin", ac.Signin())

	// route users
	e.POST("/signup", uc.Create())
	e.GET("/users", uc.GetAll())
	e.GET("/users/:id", uc.GetById())
	e.PUT("/users/:id", uc.Update(), _middleware.JWTMiddleware())
	e.DELETE("/users/:id", uc.Delete(), _middleware.JWTMiddleware())
}

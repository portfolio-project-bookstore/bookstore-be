package auth

import (
	_controller "bookstore/delivery/controller"
	_entity "bookstore/entity"
	_auth "bookstore/repository/auth"
	_utility "bookstore/utility"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	auth_repo _auth.AuthInterface
}

func New(auth _auth.AuthInterface) *AuthController {
	return &AuthController{auth_repo: auth}
}

func (ac AuthController) Signin() echo.HandlerFunc {
	return func(c echo.Context) error {
		data_signin := _entity.User{}

		if err := c.Bind(&data_signin); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid data request")
		}

		if err := _utility.SigninValidate(data_signin); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "data can't be empty")
		}

		token, err := ac.auth_repo.Signin(data_signin)
		if err != nil || token == "" {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "signin failed")
		}
		return _controller.SuccessWithDataResponse(c, "signin success", token)
	}
}

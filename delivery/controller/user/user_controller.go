package user

import (
	_controller "bookstore/delivery/controller"
	_middleware "bookstore/delivery/middleware"
	_entity "bookstore/entity"
	_user "bookstore/repository/user"
	_utility "bookstore/utility"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	user_repo _user.UserInterface
}

func New(user _user.UserInterface) *UserController {
	return &UserController{user_repo: user}
}

func (uc UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		new_user := _entity.User{}

		if err := c.Bind(&new_user); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid data request")
		}
		if err := _utility.UserValidate(new_user); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "data can't be empty")
		}
		new_user.Password, _ = _utility.HashPassword(new_user.Password)
		if _, err := uc.user_repo.Create(new_user); err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "email or password already exist")
		}
		return _controller.SuccessNonDataResponse(c, "success operation")
	}
}

func (uc UserController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uc.user_repo.GetAll()
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to get data")
		} else if len(users) == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "data not found")
		}
		var user_response []UserResponseFormat
		for _, value := range users {
			user_response = append(user_response, FormattingUserResponse(value))
		}
		return _controller.SuccessWithDataResponse(c, "success operation", user_response)
	}
}

func (uc UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		}

		user, len_user, err := uc.user_repo.GetById(id)
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to get data")
		} else if len_user == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "data not found")
		}

		user_response := FormattingUserResponse(user)
		return _controller.SuccessWithDataResponse(c, "success operation", user_response)
	}
}

func (uc UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		if valid := _middleware.ValidateToken(c); !valid {
			return _controller.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		}
		user, len_user, err := uc.user_repo.GetById(id)
		if id != int(user.ID) {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "access forbidden")
		}
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to get data")
		} else if len_user == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "data not found")
		}
		update_user := _entity.User{}
		if err := c.Bind(&update_user); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid data request")
		}
		if err := _utility.UserValidate(update_user); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "data can't be empty")
		}
		update_user.Password, _ = _utility.HashPassword(update_user.Password)
		if _, err := uc.user_repo.Update(id, update_user); err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "email or password already exist")
		}
		return _controller.SuccessNonDataResponse(c, "success operation")
	}
}

func (uc UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		if valid := _middleware.ValidateToken(c); !valid {
			return _controller.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		}
		user, len_user, err := uc.user_repo.GetById(id)
		if id != int(user.ID) {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "access forbidden")
		}
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to get data")
		} else if len_user == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "data not found")
		}
		if err := uc.user_repo.Delete(id); err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to delete data")
		}
		return _controller.SuccessNonDataResponse(c, "success operation")
	}
}

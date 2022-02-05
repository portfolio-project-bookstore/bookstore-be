package utility

import (
	_entity "bookstore/entity"

	"github.com/go-playground/validator/v10"
)

type UserValidator struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Address  string `validate:"required"`
	Password string `validate:"required"`
	Role     string `validate:"required"`
}

func UserValidate(user _entity.User) error {
	v := validator.New()
	user_validate := UserValidator{
		Username: user.Username,
		Email:    user.Email,
		Address:  user.Address,
		Password: user.Password,
	}
	err := v.Struct(user_validate)
	if err != nil {
		return err
	}
	return nil
}

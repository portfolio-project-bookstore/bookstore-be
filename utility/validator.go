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
		Role:     user.Role,
	}

	if err := v.Struct(user_validate); err != nil {
		return err
	}
	return nil
}

type SigninValidator struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

func SigninValidate(data_signin _entity.User) error {
	v := validator.New()
	signin_validate := SigninValidator{
		Username: data_signin.Username,
		Password: data_signin.Password,
	}

	if err := v.Struct(signin_validate); err != nil {
		return err
	}
	return nil
}

type BookValidator struct {
	Title       string `validate:"required"`
	Category    string `validate:"required"`
	Author      string `validate:"required"`
	Description string `validate:"required"`
	Price       int    `validate:"required,gte=1"`
	Stock       int    `validate:"required,gte=1"`
}

func BookValidate(book _entity.Book) error {
	v := validator.New()
	book_validate := BookValidator{
		Title:       book.Title,
		Category:    book.Category,
		Author:      book.Author,
		Description: book.Description,
		Price:       book.Price,
		Stock:       book.Stock,
	}

	if err := v.Struct(book_validate); err != nil {
		return err
	}
	return nil
}

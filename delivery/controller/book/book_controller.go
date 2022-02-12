package book

import (
	_controller "bookstore/delivery/controller"
	_middleware "bookstore/delivery/middleware"
	_entity "bookstore/entity"
	_book "bookstore/repository/book"
	_utility "bookstore/utility"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	book_repo _book.BookInterface
}

func New(book _book.BookInterface) *BookController {
	return &BookController{book_repo: book}
}

func (bc BookController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		new_book := _entity.Book{}
		if err := c.Bind(&new_book); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid data request")
		}
		if err := _utility.BookValidate(new_book); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "data can't be empty")
		}
		id_user, role := _middleware.ExtractTokenId(c)
		fmt.Println("role", role)
		fmt.Println("id_user", id_user)
		new_book.UsersID = id_user
		if role != "supplier" {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "you are not supplier")
		}
		if _, err := bc.book_repo.Create(new_book); err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to create data")
		}
		return _controller.SuccessNonDataResponse(c, "success operation")
	}
}

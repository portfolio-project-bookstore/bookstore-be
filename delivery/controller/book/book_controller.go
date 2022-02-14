package book

import (
	_controller "bookstore/delivery/controller"
	_middleware "bookstore/delivery/middleware"
	_entity "bookstore/entity"
	_book "bookstore/repository/book"
	_utility "bookstore/utility"
	"net/http"
	"strconv"

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

func (bc BookController) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := bc.book_repo.GetAll()
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to get data")
		} else if len(books) == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "data not found")
		}
		var book_response []BookResponseFormat
		for _, value := range books {
			book_response = append(book_response, FormattingBookResponse(value))
		}
		return _controller.SuccessWithDataResponse(c, "success operation", book_response)
	}
}

func (bc BookController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		}

		book, len_book, err := bc.book_repo.GetById(id)
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to get data")
		} else if len_book == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "data not found")
		}

		book_response := FormattingBookResponse(book)
		return _controller.SuccessWithDataResponse(c, "success operation", book_response)
	}
}

func (bc BookController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		if valid := _middleware.ValidateToken(c); !valid {
			return _controller.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid id")
		}

		// cek apakah buku yang diupdate ada atau tidak
		book, len_book, err := bc.book_repo.GetById(id)
		if err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "failed to get data")
		} else if len_book == 0 {
			return _controller.ErrorResponse(c, http.StatusNotFound, "data not found")
		}

		// cek user yang akan mengupdate buku adalah supplier dan pemilik buku
		id_user, role := _middleware.ExtractTokenId(c)
		if book.UsersID != id_user || role != "supplier" {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "access forbidden")
		}

		// update buku
		update_book := _entity.Book{}
		if err := c.Bind(&update_book); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "invalid data request")
		}
		if err := _utility.BookValidate(update_book); err != nil {
			return _controller.ErrorResponse(c, http.StatusBadRequest, "data can't be empty")
		}
		if _, err := bc.book_repo.Update(id, update_book); err != nil {
			return _controller.ErrorResponse(c, http.StatusInternalServerError, "email or password already exist")
		}
		return _controller.SuccessNonDataResponse(c, "success operation")
	}
}

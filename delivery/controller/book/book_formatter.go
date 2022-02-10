package book

import _entity "bookstore/entity"

type CreateBookRequest struct {
	Title       string `json:"title" form:"title"`
	Category    string `json:"category" form:"category"`
	Author      string `json:"author" form:"author"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	Stock       int    `json:"stock" form:"stock"`
}

type BookResponseFormat struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Category    string `json:"category"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}

func FormattingBookResponse(format _entity.Book) BookResponseFormat {
	return BookResponseFormat{
		ID:          format.ID,
		Title:       format.Title,
		Category:    format.Category,
		Author:      format.Author,
		Description: format.Description,
		Price:       format.Price,
		Stock:       format.Stock,
	}
}

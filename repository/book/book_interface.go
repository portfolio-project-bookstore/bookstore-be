package book

import _entity "bookstore/entity"

type BookInterface interface {
	Create(new_book _entity.Book) (_entity.Book, error)
	GetAll() ([]_entity.Book, error)
	GetById(id int) (_entity.Book, int, error)
	Update(id int, update_book _entity.Book) (_entity.Book, error)
	Delete(id int) error
}

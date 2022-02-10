package book

import (
	_entity "bookstore/entity"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (br *BookRepository) Create(new_book _entity.Book) (_entity.Book, error) {
	if err := br.db.Save(&new_book).Error; err != nil {
		return new_book, err
	}
	return new_book, nil
}

func (br *BookRepository) GetAll() ([]_entity.Book, error) {
	books := []_entity.Book{}
	if err := br.db.Table("books").Select("*").Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func (br *BookRepository) GetById(id int) (_entity.Book, int, error) {
	book := _entity.Book{}
	query := br.db.Table("books").Select("*").Where("books.id = ?", id).Find(&book)
	if query.Error != nil || query.RowsAffected == 0 {
		return book, int(query.RowsAffected), query.Error
	}
	return book, int(query.RowsAffected), nil
}

func (br *BookRepository) Update(id int, update_book _entity.Book) (_entity.Book, error) {
	if err := br.db.Where("id = ?", id).Updates(&update_book).Error; err != nil {
		return update_book, err
	}
	br.db.First(&update_book, id)
	return update_book, nil
}

func (br *BookRepository) Delete(id int) error {
	var book _entity.Book
	if err := br.db.Where("id = ?", id).Delete(&book).Error; err != nil {
		return err
	}
	return nil
}

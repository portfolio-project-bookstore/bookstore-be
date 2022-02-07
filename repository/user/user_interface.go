package user

import _entity "bookstore/entity"

type UserInterface interface {
	Create(new_user _entity.User) (_entity.User, error)
	GetAll() ([]_entity.User, error)
	GetById(id int) (_entity.User, int, error)
	Update(id int, update_user _entity.User) (_entity.User, error)
	Delete(id int) error
}

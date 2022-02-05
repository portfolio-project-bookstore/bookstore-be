package user

import (
	_entity "bookstore/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(new_user _entity.User) (_entity.User, error) {
	if err := ur.db.Save(&new_user).Error; err != nil {
		return new_user, err
	}
	return new_user, nil
}

func (ur *UserRepository) GetAll() ([]_entity.User, error) {
	users := []_entity.User{}
	if err := ur.db.Table("users").Select("*").Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}

func (ur *UserRepository) GetById(id int) (_entity.User, int, error) {
	user := _entity.User{}
	query := ur.db.Table("users").Select("*").Where("users.id = ?", id).Find(&user)
	if query.Error != nil || query.RowsAffected == 0 {
		return user, int(query.RowsAffected), query.Error
	}
	return user, int(query.RowsAffected), nil
}

func (ur *UserRepository) Update(id int, update_user _entity.User) (_entity.User, error) {
	if err := ur.db.Where("id = ?", id).Updates(&update_user).Error; err != nil {
		return update_user, err
	}
	ur.db.First(&update_user, id)
	return update_user, nil
}

func (ur *UserRepository) Delete(id int) error {
	var user _entity.User
	if err := ur.db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

package auth

import (
	_middleware "bookstore/delivery/middleware"
	_entity "bookstore/entity"
	_utility "bookstore/utility"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) Signin(data_signin _entity.User) (string, error) {
	var user _entity.User
	if err := ar.db.Where("username = ?", data_signin.Username).First(&user).Error; err != nil {
		return "", err
	}

	if match := _utility.DecryptPassword(user.Password, data_signin.Password); !match {
		return "", nil
	}

	token, err := _middleware.CreateToken(int(user.ID), user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

package auth

import _entity "bookstore/entity"

type AuthInterface interface {
	Signin(data_signin _entity.User) (string, error)
}

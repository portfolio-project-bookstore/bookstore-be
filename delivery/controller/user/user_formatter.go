package user

import _entity "bookstore/entity"

type CreateUserRequest struct {
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

type UserResponseFormat struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Role     string `json:"role"`
}

func FormattingUserResponse(format _entity.User) UserResponseFormat {
	return UserResponseFormat{
		ID:       format.ID,
		Username: format.Username,
		Email:    format.Email,
		Address:  format.Address,
		Role:     format.Role,
	}
}

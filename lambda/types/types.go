package types

import (
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username     string
	PasswordHash string
}

func NewUser(req RegisterUserRequest) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	return User{
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
	}, nil
}

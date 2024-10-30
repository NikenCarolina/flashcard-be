package model

import "github.com/NikenCarolina/flashcard-be/internal/dto"

type User struct {
	UserID       int
	Username     string
	PasswordHash string
}

func (u *User) ToDto() *dto.User {
	return &dto.User{
		Username: u.Username,
	}
}

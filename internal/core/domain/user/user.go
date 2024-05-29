package user

import "github.com/google/uuid"

type User struct {
	Id    uuid.UUID
	Email Email
}

func New(email Email) *User {
	return &User{
		Id:    uuid.New(),
		Email: email,
	}
}

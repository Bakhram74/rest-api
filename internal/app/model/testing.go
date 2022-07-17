package model

import "testing"

func TestingUser(t *testing.T) *User {
	return &User{
		Email:    "example@mail.com",
		Password: "password",
	}
}

package model

import "testing"

func TestingUser(t *testing.T) *User {
	return &User{
		Email:    "example@list.com",
		Password: "password",
	}
}

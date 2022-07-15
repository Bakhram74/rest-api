package model_test

import (
	"github.com/Bakhram74/rest-api.git/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_Validation(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestingUser(t)
			},
			isValid: true,
		},
		{
			name: "email is empty",
			u: func() *model.User {
				u := model.TestingUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestingUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestingUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestingUser(t)
				u.Password = "short1"
				return u
			},
			isValid: false,
		},
		{
			name: "encrypted password",
			u: func() *model.User {
				u := model.TestingUser(t)
				u.Password = ""
				u.EncryptedPassword = "not nil"
				return u
			},
			isValid: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validation())
			} else {
				assert.Error(t, tc.u().Validation())
			}
		})
	}
}
func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestingUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

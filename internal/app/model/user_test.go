package model_test

import (
	"github.com/Bakhram74/rest-api.git/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestingUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}

package store_test

import (
	"github.com/Bakhram74/rest-api.git/internal/app/model"
	"github.com/Bakhram74/rest-api.git/internal/app/store"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")
	email := "example@.com"
	_, err := s.User().FindByEmail(email)
	if err != nil {
		return
	}
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "example@.com",
	})

	u, err := s.User().FindByEmail(email)
	if err != nil {
		return
	}
	assert.NotNil(t, u)
	assert.NoError(t, err)
}

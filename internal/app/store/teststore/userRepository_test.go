package teststore_test

import (
	"github.com/Bakhram74/rest-api.git/internal/app/model"
	"github.com/Bakhram74/rest-api.git/internal/app/store"
	"github.com/Bakhram74/rest-api.git/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	err := s.User().Create(model.TestingUser(t))
	u := model.TestingUser(t)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	u := model.TestingUser(t)
	s.User().Create(u)
	u2, err := s.User().Find(u.ID)

	assert.NotNil(t, u2)
	assert.NoError(t, err)
}
func TestUserRepository_FindByEmail(t *testing.T) {
	email := "example@mail.com"
	s := teststore.New()
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestingUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)

	assert.NotNil(t, u)
	assert.NoError(t, err)
}

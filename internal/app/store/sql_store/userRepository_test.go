package sql_store_test

import (
	"github.com/Bakhram74/rest-api.git/internal/app/model"
	"github.com/Bakhram74/rest-api.git/internal/app/store"
	"github.com/Bakhram74/rest-api.git/internal/app/store/sql_store"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sql_store.TestDB(t, databaseUrl)
	defer teardown("users")
	s := sql_store.New(db)
	err := s.User().Create(model.TestingUser(t))
	u := model.TestingUser(t)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sql_store.TestDB(t, databaseUrl)
	defer teardown("users")
	email := "example@mail.com"
	s := sql_store.New(db)
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestingUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)

	assert.NotNil(t, u)
	assert.NoError(t, err)
}

package sql_store

import (
	"database/sql"
	"github.com/Bakhram74/rest-api.git/internal/app/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) User() store.UserRepository {
	// todo
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}

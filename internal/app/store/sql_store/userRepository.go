package sql_store

import (
	"database/sql"
	"github.com/Bakhram74/rest-api.git/internal/app/model"
	"github.com/Bakhram74/rest-api.git/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {
	err := u.Validation()
	if err != nil {
		return err
	}
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	return r.store.db.QueryRow(
		"INSERT INTO users (email,encrypted_password) VALUES ($1,$2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)

}
func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT id,email,encrypted_password FROM users WHERE id=$1",
		id,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword)
	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	err := r.store.db.QueryRow(
		"SELECT id,email,encrypted_password FROM users WHERE email=$1",
		email,
	).Scan(&u.ID, &u.Email, &u.EncryptedPassword)
	if err == sql.ErrNoRows {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}

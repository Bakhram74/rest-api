package store

import "github.com/Bakhram74/rest-api.git/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	err := r.store.db.QueryRow(
		"INSERT INTO users (email,encrypted_password) VALUES (&1,&2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}

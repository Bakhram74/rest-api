package store

import "github.com/Bakhram74/rest-api.git/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}

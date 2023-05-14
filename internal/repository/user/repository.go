package user

import "github.com/jesslyn-ctrl/go-restaurant-app/internal/model"

type Repository interface {
	RegisterUser(userData model.User) (model.User, error)
	CheckRegistered(username string) (bool, error)
	GenerateUserHash(password string) (hash string, err error)
}

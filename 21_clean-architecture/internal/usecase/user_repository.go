package usecase

import "project/internal/domain/model"

type UserRepository interface {
	GetUserByID(id int64) (*model.User, error)
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
}

package usecase

import "project/internal/domain/model"

type UserRepository interface {
	GetUserByID(id int64) (*model.User, error)
	CreateUser(user *model.User) error
}

type UserUseCase struct {
	UserRepo UserRepository
}

func (u *UserUseCase) CreateUser(user *model.User) error {
	return u.UserRepo.CreateUser(user)
}

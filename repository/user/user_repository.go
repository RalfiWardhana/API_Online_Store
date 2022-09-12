package user

import "project/online-store/entity"

type UserRepository interface {
	FindAllUsers() ([]entity.User, error)
	FindUserById(id string) (entity.User, error)
	UpdateUserById(id string, user entity.User) error
	DeleteUserById(id string, user entity.User) error
}

package auth

import "project/online-store/entity"

type AuthRepository interface {
	Register(user entity.User) (entity.User, error)
	Login(email string) (entity.User, error)

	ActivatedStore(id string, role int) error
}

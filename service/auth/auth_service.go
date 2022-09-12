package auth

import "project/online-store/entity"

type AuthService interface {
	Register(user entity.User) (entity.User, error)
	Login(email, password string) (string, error)

	ActivatedStore(id string, role int) error
}

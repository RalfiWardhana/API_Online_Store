package auth

import (
	"project/online-store/config"
	"project/online-store/entity"
	"project/online-store/repository/auth"
	"project/online-store/utils"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	r    auth.AuthRepository
	conf config.Config
}

func NewAuthService(r auth.AuthRepository, conf config.Config) AuthService {
	return &Service{
		r:    r,
		conf: conf,
	}
}

func (s *Service) Register(user entity.User) (entity.User, error) {
	// hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hash)
	// return user
	return s.r.Register(user)
}

func (s *Service) Login(email, password string) (string, error) {
	// get user
	user, er := s.r.Login(email)
	if er != nil {
		return "", er
	}

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	// generate token
	token, _ := utils.GenerateAccessToken(user.Id, user.Email, s.conf.JWT_SECRET, user.RoleId)

	return token, nil
}

func (s *Service) ActivatedStore(id string, role int) error {
	// update role user to admin
	return s.r.ActivatedStore(id, role)
}

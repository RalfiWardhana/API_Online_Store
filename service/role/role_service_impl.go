package role

import (
	"project/online-store/config"
	"project/online-store/entity"
	"project/online-store/repository/role"
)

type Service struct {
	r    role.RoleRepository
	conf config.Config
}

func NewRoleService(r role.RoleRepository, conf config.Config) RoleService {
	return &Service{
		r:    r,
		conf: conf,
	}
}

func (s *Service) CreateRole(role entity.Role) error {
	err := s.r.CreateRole(role)
	if err != nil {
		return err
	}

	return nil
}

package role

import "project/online-store/entity"

type RoleService interface {
	CreateRole(role entity.Role) error
}

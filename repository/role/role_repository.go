package role

import "project/online-store/entity"

type RoleRepository interface {
	CreateRole(role entity.Role) error
}

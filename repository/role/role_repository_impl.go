package role

import (
	"project/online-store/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateRole(role entity.Role) error {
	err := r.db.Create(&role).Error
	if err != nil {
		return err
	}

	return nil
}

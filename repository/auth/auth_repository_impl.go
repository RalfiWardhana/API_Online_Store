package auth

import (
	"project/online-store/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Register(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Login(email string) (entity.User, error) {
	user := entity.User{}
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) ActivatedStore(id string, role int) error {
	user := entity.User{}
	// update user role to admin
	err := r.db.Model(&user).Where("id = ?", id).Update("role_id", role).Error
	if err != nil {
		return err
	}
	return nil
}

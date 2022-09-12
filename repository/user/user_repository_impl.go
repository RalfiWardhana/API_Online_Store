package user

import (
	"project/online-store/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAllUsers() ([]entity.User, error) {
	user := []entity.User{}
	err := r.db.Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindUserById(id string) (entity.User, error) {
	user := entity.User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) UpdateUserById(id string, user entity.User) error {
	err := r.db.Model(&user).Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteUserById(id string, user entity.User) error {
	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

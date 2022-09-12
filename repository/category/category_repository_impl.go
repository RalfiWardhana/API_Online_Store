package category

import (
	"project/online-store/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateCategory(category entity.Category) error {
	err := r.db.Create(&category).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAllCategory() ([]entity.Category, error) {
	category := []entity.Category{}
	err := r.db.Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindCategoryByID(id string) (entity.Category, error) {
	category := entity.Category{}
	err := r.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

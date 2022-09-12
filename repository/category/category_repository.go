package category

import "project/online-store/entity"

type CategoryRepository interface {
	CreateCategory(category entity.Category) error
	FindAllCategory() ([]entity.Category, error)
	FindCategoryByID(id string) (entity.Category, error)
}

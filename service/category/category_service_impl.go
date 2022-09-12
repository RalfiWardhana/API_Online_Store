package category

import (
	"project/online-store/config"
	"project/online-store/entity"
	"project/online-store/repository/category"
)

type Service struct {
	r    category.CategoryRepository
	conf config.Config
}

func NewCategoryService(r category.CategoryRepository, conf config.Config) CategoryService {
	return &Service{
		r:    r,
		conf: conf,
	}
}

func (s *Service) CreateCategory(category entity.Category) error {
	err := s.r.CreateCategory(category)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FindAllCategory() ([]entity.Category, error) {
	category, err := s.r.FindAllCategory()
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s *Service) FindCategoryByID(id string) (entity.Category, error) {
	category, err := s.r.FindCategoryByID(id)
	if err != nil {
		return category, err
	}

	return category, nil
}

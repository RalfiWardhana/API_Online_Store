package products

import (
	"project/online-store/entity"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) FindAllProducts() ([]entity.Product, error) {
	product := []entity.Product{}

	err := r.db.Model(&entity.Product{}).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) CreateProduct(product entity.Product) error {
	err := r.db.Create(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindProductByID(id string) (entity.Product, error) {
	product := entity.Product{}
	err := r.db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) UpdateProduct(id string, product entity.Product) (entity.Product, error) {
	err := r.db.Where("id = ?", id).Updates(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) DeleteProduct(id string) error {
	product := entity.Product{}
	err := r.db.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}

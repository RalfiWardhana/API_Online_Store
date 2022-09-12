package products

import "project/online-store/entity"

type ProductRepository interface {
	FindAllProducts() ([]entity.Product, error)
	FindProductByID(id string) (entity.Product, error)
	CreateProduct(product entity.Product) error
	UpdateProduct(id string, product entity.Product) (entity.Product, error)
	DeleteProduct(id string) error
}

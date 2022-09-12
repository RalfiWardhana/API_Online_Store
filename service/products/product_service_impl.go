package products

import (
	"project/online-store/config"
	"project/online-store/entity"
	"project/online-store/repository/products"
)

type Service struct {
	r    products.ProductRepository
	conf config.Config
}

func NewProductService(r products.ProductRepository, conf config.Config) ProductService {
	return &Service{
		r:    r,
		conf: conf,
	}
}

func (s *Service) FindAllProducts() ([]entity.Product, error) {
	product, err := s.r.FindAllProducts()
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *Service) CreateProduct(product entity.Product) error {
	err := s.r.CreateProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FindProductByID(id string) (entity.Product, error) {
	product, err := s.r.FindProductByID(id)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *Service) UpdateProduct(id string, products entity.Product) (entity.Product, error) {
	_, er := s.r.FindProductByID(id)
	if er != nil {
		return products, er
	}
	product, err := s.r.UpdateProduct(id, products)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *Service) DeleteProduct(id string) error {
	_, err := s.r.FindProductByID(id)
	if err != nil {
		return err
	}
	er := s.r.DeleteProduct(id)
	if er != nil {
		return er
	}

	return nil
}

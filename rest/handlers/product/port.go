package product

import "ecommerce/domain"

type Service interface {
	Create(domain.Product) (*domain.Product, error)
	Update(domain.Product) (*domain.Product, error)
	Get(int) (*domain.Product, error)
	List() ([]*domain.Product, error)
	Delete(int) error
}
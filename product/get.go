package product

import "ecommerce/domain"

func (svc *service) Get(productId int) (*domain.Product, error) {
	return svc.prdctRepo.Get(productId)
}

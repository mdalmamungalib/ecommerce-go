package user

import "ecommerce/domain"

type service struct {
	usrRepo UserREpo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.usrRepo.Create(user)
	
}
func (svc *service) Find(email string, pass string) (*domain.User, error)
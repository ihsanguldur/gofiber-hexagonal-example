package services

import "weatherApp/domain"

type service struct {
	userRepo domain.Repository
}

func NewUserService(userRepo domain.Repository) *service {
	return &service{userRepo: userRepo}
}

func (s *service) Create(user *domain.User) error {
	return s.userRepo.Create(user)
}

package service

import (
	models "kilosale_main/internal/models/user"
	repository "kilosale_main/internal/repositories/user"
)

type UserService interface {
	Create(user *models.User) error
	GetByPhone(phone string) (*models.User, error)
	GetByID(id string) (*models.User, error)
	Update(user *models.User) error
	Delete(user *models.User) error
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

type userService struct {
	repo repository.UserRepository
}

func (s *userService) Create(user *models.User) error {
	return s.repo.Create(user)
}

func (s *userService) GetByPhone(phone string) (*models.User, error) {
	return s.repo.GetByPhone(phone)
}

func (s *userService) GetByID(id string) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) Update(user *models.User) error {
	return s.repo.Update(user)
}

func (s *userService) Delete(user *models.User) error {
	return s.repo.Delete(user)
}

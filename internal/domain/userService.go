package domain

import (
	"host/internal/domain/interfaces"
	"host/internal/domain/models"
)

type UserService struct {
	repo interfaces.IUserRepo
}

func NewUserService(repo interfaces.IUserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(userCreationParams *models.CreateUserRequest) error {
	user := &models.User{
		Name:    userCreationParams.Name,
		Profile: userCreationParams.Profile,
	}

	return s.repo.Create(user)
}

func (s *UserService) Update(user *models.User) error {
	return s.repo.Update(user)
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}

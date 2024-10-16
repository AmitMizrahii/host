package interfaces

import "host/internal/domain/models"

type IUserService interface {
	Create(user *models.CreateUserRequest) error
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

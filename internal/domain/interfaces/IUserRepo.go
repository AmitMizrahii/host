package interfaces

import "host/internal/domain/models"

type IUserRepo interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

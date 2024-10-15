package dal

import (
	"host/types"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *types.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetByID(id uint) (*types.User, error) {
	var user types.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) Update(user *types.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&types.User{}, id).Error
}

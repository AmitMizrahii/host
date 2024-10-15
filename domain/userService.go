package domain

import "host/types"

type UserService struct {
	repo types.IUserRepo
}

func NewUserService(repo types.IUserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(userCraetionParams *types.CreateUserRequest) error {
	user := &types.User{
		Name:    userCraetionParams.Name,
		Profile: userCraetionParams.Profile,
	}

	return s.repo.Create(user)
}

func (s *UserService) Update(user *types.User) error {
	return s.repo.Update(user)
}

func (s *UserService) GetByID(id uint) (*types.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Delete(id uint) error {
	return s.repo.Delete(id)
}

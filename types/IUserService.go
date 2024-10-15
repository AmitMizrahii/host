package types

type IUserService interface {
	Create(user *CreateUserRequest) error
	GetByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

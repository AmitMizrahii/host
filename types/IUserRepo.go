package types

type IUserRepo interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

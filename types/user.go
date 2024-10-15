package types

type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

type CreateUserRequest struct {
	Name    string `json:"name" binding:"required,min=2"`
	Profile string `json:"profile" binding:"required"`
}

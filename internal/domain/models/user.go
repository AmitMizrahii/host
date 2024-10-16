package models

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

type CreateUserRequest struct {
	Name    string `json:"name" validate:"required,min=2"`
	Profile string `json:"profile" validate:"required"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *CreateUserRequest) Validate() error {
	return validate.Struct(u)
}

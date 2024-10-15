package client

import "host/types"

type Client struct {
	service types.IUserService
}

func NewUserController(service types.IUserService) *Client {
	return &Client{service: service}
}

func (c *Client) CreateUser(userCraetionParams *types.CreateUserRequest) error {
	return c.service.Create(userCraetionParams)
}

func (c *Client) UpdateUser(user *types.User) error {
	return c.service.Update(user)
}

func (c *Client) GetUserByID(id uint) (*types.User, error) {
	return c.service.GetByID(id)
}

func (c *Client) Delete(id uint) error {
	return c.service.Delete(id)
}

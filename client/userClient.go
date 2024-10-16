package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"host/internal/domain/models"
	"io"

	"net/http"
)

type UserClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewUserClient(baseURL string) *UserClient {
	return &UserClient{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *UserClient) CreateUser(user models.CreateUserRequest) (*models.User, error) {
	url := fmt.Sprintf("%s/users", c.BaseURL)

	responseBody, err := c.sendRequest(http.MethodPost, url, user)
	if err != nil {
		return nil, fmt.Errorf("create user failed: %w", err)
	}

	var createdUser models.User
	if err := json.Unmarshal(responseBody, &createdUser); err != nil {
		return nil, fmt.Errorf("failed to parse create user response: %w", err)
	}

	return &createdUser, nil
}

func (c *UserClient) GetUser(userID uint) (*models.User, error) {
	url := fmt.Sprintf("%s/users/%d", c.BaseURL, userID)

	responseBody, err := c.sendRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("get user failed: %w", err)
	}

	var user models.User
	if err := json.Unmarshal(responseBody, &user); err != nil {
		return nil, fmt.Errorf("failed to parse get user response: %w", err)
	}

	return &user, nil
}

func (c *UserClient) UpdateUser(userID uint, user models.User) (*models.User, error) {
	url := fmt.Sprintf("%s/users/%d", c.BaseURL, userID)

	responseBody, err := c.sendRequest(http.MethodPut, url, user)
	if err != nil {
		return nil, fmt.Errorf("update user failed: %w", err)
	}

	var updatedUser models.User
	if err := json.Unmarshal(responseBody, &updatedUser); err != nil {
		return nil, fmt.Errorf("failed to parse update user response: %w", err)
	}

	return &updatedUser, nil
}

func (c *UserClient) DeleteUser(userID uint) error {
	url := fmt.Sprintf("%s/users/%d", c.BaseURL, userID)

	_, err := c.sendRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("delete user failed: %w", err)
	}

	return nil
}

func (c *UserClient) sendRequest(method, url string, data interface{}) ([]byte, error) {
	var body []byte
	var err error

	if data != nil {
		body, err = json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("server returned error: %s", resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return responseBody, nil
}

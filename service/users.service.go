package service

import (
	"errors"
	"fmt"
	"golang-fuego-sqlite/controller"
	"golang-fuego-sqlite/models"
	"log/slog"
)

func NewInMemoryUsersService() *InMemoryUsersService {
	return &InMemoryUsersService{
		Users: []models.Users{},
		Incr:  new(int),
	}
}

type InMemoryUsersService struct {
	Users []models.Users
	Incr  *int
}

// CreateUsers implements controller.UsersService.
func (userService *InMemoryUsersService) CreateUsers(c models.UsersCreate) (models.Users, error) {
	*userService.Incr++
	newUser := models.Users{
		ID:   fmt.Sprintf("user-%d", *userService.Incr),
		Name: c.Name,
		Age:  c.Age,
	}
	userService.Users = append(userService.Users, newUser)
	slog.Info("Created user", "id", newUser.ID)

	return newUser, nil
}

// DeleteUsers implements controller.UsersService.
func (userService *InMemoryUsersService) DeleteUsers(id string) (any, error) {
	for i, p := range userService.Users {
		if p.ID == id {
			userService.Users = append(userService.Users[:i], userService.Users[i+1:]...)
			return nil, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetUsers implements controller.UsersService.
func (userService *InMemoryUsersService) GetUsers(id string) (models.Users, error) {
	for _, p := range userService.Users {
		if p.ID == id {
			return p, nil
		}
	}
	return models.Users{}, errors.New("user not found")
}

var _ controller.UsersService = &InMemoryUsersService{}

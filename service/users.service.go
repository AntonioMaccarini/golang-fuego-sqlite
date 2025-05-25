package service

import (
	"errors"
	"golang-fuego-sqlite/controller"
	"golang-fuego-sqlite/models"
	"log/slog"
)

func NewInMemoryUsersService() *InMemoryUsersService {
	return &InMemoryUsersService{
		Users: []models.User{},
		Incr:  new(uint),
	}
}

type InMemoryUsersService struct {
	Users []models.User
	Incr  *uint
}

// CreateUsers implements controller.UsersService.
func (userService *InMemoryUsersService) CreateUsers(c models.UsersCreate) (models.User, error) {
	*userService.Incr++
	newUser := models.User{
		ID:   *userService.Incr,
		Name: c.Name,
	}
	userService.Users = append(userService.Users, newUser)
	slog.Info("Created user", "id", newUser.ID)

	return newUser, nil
}

// DeleteUsers implements controller.UsersService.
func (userService *InMemoryUsersService) DeleteUsers(id uint) (any, error) {
	for i, p := range userService.Users {
		if p.ID == id {
			userService.Users = append(userService.Users[:i], userService.Users[i+1:]...)
			return nil, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetUsers implements controller.UsersService.
func (userService *InMemoryUsersService) GetUsers(id uint) (models.User, error) {
	for _, p := range userService.Users {
		if p.ID == id {
			return p, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

var _ controller.UsersService = &InMemoryUsersService{}

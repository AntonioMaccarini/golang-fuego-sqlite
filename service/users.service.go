package service

import (
	"golang-fuego-sqlite/controller"
	"golang-fuego-sqlite/models"
	"golang-fuego-sqlite/queries"
)

type UsersServiceImpl struct {
	Repo *queries.UserQueries
}

func (s *UsersServiceImpl) CreateUsers(input models.UsersCreate) (models.User, error) {
	user := models.User{
		Name:  input.Name,
		Email: input.Email,
	}
	createdUser, err := s.Repo.CreateUser(&user)
	return *createdUser, err
}

func (s *UsersServiceImpl) GetUsers(id uint) (models.User, error) {
	user, err := s.Repo.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}
	return *user, nil
}

func (s *UsersServiceImpl) DeleteUsers(id uint) (any, error) {
	return nil, s.Repo.DeleteUser(id)
}

var _ controller.UsersService = &UsersServiceImpl{}

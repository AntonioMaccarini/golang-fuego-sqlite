package queries

import (
	"errors"
	"golang-fuego-sqlite/models"

	"github.com/go-fuego/fuego"
	"gorm.io/gorm"
)

type UserQueries struct {
	DB *gorm.DB
}

func (q *UserQueries) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := q.DB.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fuego.NotFoundError{
			Title:  "User not found",
			Detail: "No user with this ID",
			Err:    err,
		}
	}
	return &user, err
}

func (q *UserQueries) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := q.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (q *UserQueries) GetUsers() ([]models.User, error) {
	var users []models.User
	err := q.DB.Find(&users).Error
	return users, err
}

func (q *UserQueries) CreateUser(user *models.User) (*models.User, error) {
	err := q.DB.Create(user).Error
	if err != nil {
		return nil, fuego.InternalServerError{
			Detail: "Failed to create user",
			Err:    err,
		}
	}
	return user, nil
}

func (q *UserQueries) UpdateUser(user *models.User) (*models.User, error) {
	err := q.DB.Save(user).Error
	return user, err
}

func (q *UserQueries) DeleteUser(id uint) error {
	return q.DB.Delete(&models.User{}, id).Error
}

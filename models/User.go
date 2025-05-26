package models

import (
	"context"
	"errors"

	"github.com/go-fuego/fuego"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name" validate:"required" example:"Napoleon" gorm:"not null"`
	Email string `json:"email" validate:"required,email" example:"napoleon@example.com" gorm:"unique;not null"`
}

type UsersCreate struct {
	Name  string `json:"name" validate:"required,min=1,max=100"`
	Email string `json:"email" validate:"required,email"`
}

type UsersUpdate struct {
	Name string `json:"name,omitempty" validate:"min=1,max=100"`
}

var _ fuego.InTransformer = &User{}

func (*User) InTransform(context.Context) error {
	return errors.New("users must only be used as output")
}

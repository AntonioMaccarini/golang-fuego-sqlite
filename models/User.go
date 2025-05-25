package models

import (
	"context"
	"errors"

	"github.com/go-fuego/fuego"
	"gorm.io/gorm"
)

// Example of generic response
type BareSuccessResponse[Res any] struct {
	StatusCode int    `json:"statusCode"`
	Result     Res    `json:"result"`
	Message    string `json:"message"`
}

type User struct {
	gorm.Model
	ID   uint   `json:"id" validate:"required" gorm:"not null; primaryKey"`
	Name string `json:"name" validate:"required" example:"Napoleon" gorm:"not null"`
}

type UsersCreate struct {
	Name string `json:"name" validate:"required,min=1,max=100" example:"Napoleon"`
}

type UsersUpdate struct {
	Name string `json:"name,omitempty" validate:"min=1,max=100" example:"Napoleon" description:"Name of the user"`
}

var _ fuego.InTransformer = &User{}

func (*User) InTransform(context.Context) error {
	return errors.New("users must only be used as output")
}

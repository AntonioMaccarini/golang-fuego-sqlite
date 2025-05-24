package server

import (
	"golang-fuego-sqlite/controller"
	"golang-fuego-sqlite/service"

	"github.com/go-fuego/fuego"
)

func NewGrafanaStoreServer(options ...func(*fuego.Server)) *fuego.Server {
	s := fuego.NewServer(options...)

	usersResources := controller.UsersResources{
		UsersService: service.NewInMemoryUsersService(),
	}
	usersResources.Routes(s)
	return s
}

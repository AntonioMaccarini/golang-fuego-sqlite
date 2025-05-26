package lib

import (
	"golang-fuego-sqlite/controller"
	"golang-fuego-sqlite/models"
	"golang-fuego-sqlite/queries"
	"golang-fuego-sqlite/service"

	"github.com/go-fuego/fuego"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewGrafanaStoreServer(options ...func(*fuego.Server)) *fuego.Server {
	s := fuego.NewServer(options...)

	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(&models.User{})

	userQueries := &queries.UserQueries{DB: db}
	userService := &service.UsersServiceImpl{Repo: userQueries}

	usersResources := controller.UsersResources{
		UsersService: userService,
	}
	usersResources.Routes(s)

	return s
}

package main

import (
	"golang-fuego-sqlite/models"
	"golang-fuego-sqlite/teste-sqlite/handlers"
	"golang-fuego-sqlite/teste-sqlite/queries"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("error connecting to  database")
	}

	db.AutoMigrate(&models.User{})

	server := fuego.NewServer()

	userQueries := &queries.UserQueries{DB: db}
	handlers := &handlers.UserResources{UserQueries: userQueries}

	fuego.Get(server, "/", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, World!", nil
	})
	fuego.Get(server, "/users", handlers.GetUsers)
	fuego.Post(server, "/users", handlers.CreateUser, option.DefaultStatusCode(201))
	fuego.Get(server, "/users/{id}", handlers.GetUserByID)
	// fuego.Delete(server, "/users/{id}", handlers.DeleteUser, option.DefaultStatusCode(204))

	server.Run()
}

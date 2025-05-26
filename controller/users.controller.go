package controller

import (
	"golang-fuego-sqlite/models"
	"strconv"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
)

type UsersResources struct {
	UsersService UsersService
}

func (rs UsersResources) Routes(s *fuego.Server) {
	usersGroup := fuego.Group(s, "/users", option.Header("X-Header", "header description"))

	fuego.Post(usersGroup, "/", rs.postUsers,
		option.DefaultStatusCode(201),
		option.AddResponse(409, "Conflict", fuego.Response{Type: UsersError{}}),
	)

	fuego.Get(usersGroup, "/{id}", rs.getUsers,
		option.OperationID("getUser"),
		option.Path("id", "User ID", param.Example("example", "1")),
	)

	fuego.Delete(usersGroup, "/{id}", rs.deleteUsers)
}

func (rs UsersResources) postUsers(c fuego.ContextWithBody[models.UsersCreate]) (models.User, error) {
	body, err := c.Body()
	if err != nil {
		return models.User{}, err
	}
	return rs.UsersService.CreateUsers(body)
}

func (rs UsersResources) getUsers(c fuego.ContextNoBody) (models.User, error) {
	idStr := c.PathParam("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return models.User{}, fuego.BadRequestError{
			Title:  "Invalid ID",
			Detail: "ID must be a number",
			Err:    err,
		}
	}
	return rs.UsersService.GetUsers(uint(idUint))
}

func (rs UsersResources) deleteUsers(c fuego.ContextNoBody) (any, error) {
	idStr := c.PathParam("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return nil, fuego.BadRequestError{
			Title:  "Invalid ID",
			Detail: "ID must be a number",
			Err:    err,
		}
	}
	return rs.UsersService.DeleteUsers(uint(idUint))
}

type UsersError struct {
	Err     error  `json:"-" xml:"-"`
	Message string `json:"message" xml:"message"`
}

func (e UsersError) Error() string { return e.Err.Error() }

type UsersService interface {
	GetUsers(id uint) (models.User, error)
	CreateUsers(models.UsersCreate) (models.User, error)
	DeleteUsers(id uint) (any, error)
}

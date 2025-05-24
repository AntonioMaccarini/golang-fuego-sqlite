package controller

import (
	"golang-fuego-sqlite/models"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
)

type UsersResources struct {
	UsersService UsersService
}

type UsersError struct {
	Err     error  `json:"-" xml:"-"`
	Message string `json:"message" xml:"message"`
}

var _ error = UsersError{}

func (e UsersError) Error() string { return e.Err.Error() }

func (rs UsersResources) Routes(s *fuego.Server) {
	usersGroup := fuego.Group(s, "/users", option.Header("X-Header", "header description"))

	fuego.Post(usersGroup, "/", rs.postUsers,
		option.DefaultStatusCode(201),
		option.AddResponse(409, "Conflict: User with the same name already exists", fuego.Response{Type: UsersError{}}),
	)

	fuego.Get(usersGroup, "/{id}", rs.getUsers,
		option.OverrideDescription("Replace description with this sentence."),
		option.OperationID("getUser"),
		option.Path("id", "User ID", param.Example("example", "123")),
	)
	fuego.Delete(usersGroup, "/{id}", rs.deleteUsers)
}

func (rs UsersResources) postUsers(c fuego.ContextWithBody[models.UsersCreate]) (models.Users, error) {
	body, err := c.Body()
	if err != nil {
		return models.Users{}, err
	}

	return rs.UsersService.CreateUsers(body)
}

func (rs UsersResources) getUsers(c fuego.ContextNoBody) (models.Users, error) {
	id := c.PathParam("id")

	return rs.UsersService.GetUsers(id)
}

func (rs UsersResources) deleteUsers(c fuego.ContextNoBody) (any, error) {
	return rs.UsersService.DeleteUsers(c.PathParam("id"))
}

type UsersFilter struct {
	Name        string
	YoungerThan int
}

type UsersService interface {
	GetUsers(id string) (models.Users, error)
	CreateUsers(models.UsersCreate) (models.Users, error)
	DeleteUsers(id string) (any, error)
}

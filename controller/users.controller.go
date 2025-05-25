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

func (rs UsersResources) postUsers(c fuego.ContextWithBody[models.UsersCreate]) (models.User, error) {
	body, err := c.Body()
	if err != nil {
		return models.User{}, err
	}

	return rs.UsersService.CreateUsers(body)
}

func (rs UsersResources) getUsers(c fuego.ContextNoBody) (models.User, error) {
	idString := c.PathParam("id")

	idUint64, _ := strconv.ParseUint(idString, 10, 64)
	id := uint(idUint64)
	return rs.UsersService.GetUsers(id)
}

func (rs UsersResources) deleteUsers(c fuego.ContextNoBody) (any, error) {
	idStr := c.PathParam("id")
	idUint64, _ := strconv.ParseUint(idStr, 10, 64)
	id := uint(idUint64)

	return rs.UsersService.DeleteUsers(id)
}

type UsersFilter struct {
	Name        string
	YoungerThan int
}

type UsersService interface {
	GetUsers(id uint) (models.User, error)
	CreateUsers(models.UsersCreate) (models.User, error)
	DeleteUsers(id uint) (any, error)
}

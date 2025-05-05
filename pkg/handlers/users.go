package handlers

import (
	"log/slog"
	"oxid-gateway-admin-api/pkg/dtos"
	"oxid-gateway-admin-api/pkg/services"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

type UsersResources struct {
	UsersService *services.UsersService
}

func (rs UsersResources) Routes(s *fuego.Server) {
	fuego.Get(s, "/users/{username}", rs.getUser,
		option.Tags("User"),
		option.Summary("Get User By Username"),
		option.OperationID("getUser"),
		option.Security(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)
}

func (ur UsersResources) getUser(c fuego.ContextNoBody) (*dtos.User, error) {
	user_context, err := GetRequestContext(c)

	if err != nil {
		return nil, err
	}

	slog.Info("dslkfjds", "user", user_context)

	username := c.PathParam("username")

	user, err := ur.UsersService.GetUser(username)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fuego.NotFoundError{Title: "Not found", Detail: "Test not found"}
	}

	return user, nil
}

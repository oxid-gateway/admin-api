package services

import (
	"context"
	"errors"
	"log/slog"
	"oxid-gateway-admin-api/pkg/db"
	"oxid-gateway-admin-api/pkg/dtos"

	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5"
)

var usersInternalError = fuego.InternalServerError{Title: "Internal Server Error", Detail: "Upstream server side error"}
var usersNotFoundError = fuego.NotFoundError{Title: "Not found", Detail: "Upstream not found"}
var usersConflictError = fuego.ConflictError{Title: "Upstream conflict", Detail: "Configuration conflicts with another upstream"}
var missingUsersConfigError = fuego.BadRequestError{Title: "Missing Upstream Config", Detail: "Missing required parameters for upstream type"}

type UsersService struct {
	Repository *db.Queries
}

func (ts UsersService) GetUser(username string) (*dtos.User, error) {
	model, err := ts.Repository.GetUserByUsername(context.Background(), username)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			slog.Error("Failed to get user", "error", err)
			return nil, usersInternalError
		}
	}

	dto := dtos.User{
		Name: model.Name,
		Username: model.Username,
		Email: model.Email,
	}

	return &dto, nil
}

func (ts UsersService) CreateUser(body db.CreateUserParams) (*db.User, error) {
	model, err := ts.Repository.CreateUser(context.Background(), body)

	if err != nil {
		slog.Error("Failed to create user", "error", err)
		return nil, usersInternalError
	}

	return model, nil
}

func (ts UsersService) LinkUserToUpstream(user_id int32, upstream_id int32) error {
	err := ts.Repository.LinkUserToUpstream(context.Background(), db.LinkUserToUpstreamParams{
		UpstreamID: upstream_id,
		UserID: user_id,
	})

	if err != nil {
		slog.Error("Failed to create user", "error", err)
		return usersInternalError
	}

	return  nil
}

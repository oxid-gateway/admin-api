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

var upstreamInternalError = fuego.InternalServerError{Title: "Internal Server Error", Detail: "Upstream server side error"}
var upstreamNotFoundError = fuego.NotFoundError{Title: "Not found", Detail: "Upstream not found"}
var upstreamConflictError = fuego.ConflictError{Title: "Upstream conflict", Detail: "Configuration conflicts with another upstream"}
var missingUpstreamConfigError = fuego.BadRequestError{Title: "Missing Upstream Config", Detail: "Missing required parameters for upstream type"}

type UpstreamsService struct {
	Repository *db.Queries
}

func (ts UpstreamsService) GetUpstream(id int32) (*dtos.Upstream, error) {
	model, err := ts.Repository.GetUpstreamById(context.Background(), id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			slog.Error("Failed to get upstream", "error", err)
			return nil, upstreamInternalError
		}
	}

	dto := dtos.Upstream{
		ID:   model.ID,
		Name: model.Name,
	}

	return &dto, nil
}

func (ts UpstreamsService) CreateUpstream(body dtos.UpstreamCreate) (*dtos.Upstream, error) {
	conflict_model, err := ts.Repository.GetUpstreamConflic(context.Background(), db.GetUpstreamConflicParams{
		ID:   0,
		Name: body.Name,
	})

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		slog.Error("Failed to create upstream", "error", err)
		return nil, upstreamInternalError
	}

	if conflict_model.ID != 0 {
		return nil, upstreamConflictError
	}

	model, err := ts.Repository.CreateUpstream(context.Background(), body.Name)

	if err != nil {
		slog.Error("Failed to create upstream", "error", err)
		return nil, upstreamInternalError
	}

	dto := dtos.Upstream{
		ID:   model.ID,
		Name: model.Name,
	}

	return &dto, nil
}

func (ts UpstreamsService) UpdateUpstream(id int32, body dtos.UpstreamUpdate) (*dtos.Upstream, error) {
	prev_upstream, err := ts.GetUpstream(id)

	if err != nil {
		return nil, err
	}

	if prev_upstream == nil {
		return nil, nil
	}

	conflict_model, err := ts.Repository.GetUpstreamConflic(context.Background(), db.GetUpstreamConflicParams{
		ID:   id,
		Name: body.Name,
	})

	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		slog.Error("Failed to update upstream", "error", err)
		return nil, upstreamInternalError
	}

	if conflict_model.ID != 0 {
		return nil, upstreamConflictError
	}

	err = ts.Repository.UpdateUpstream(context.Background(), db.UpdateUpstreamParams{
		ID:   id,
		Name: body.Name,
	})

	if err != nil {
		slog.Error("Failed to update upstream", "error", err)
		return nil, upstreamInternalError
	}

	dto := dtos.Upstream{
		ID:   id,
		Name: body.Name,
	}

	return &dto, nil
}

func (ts UpstreamsService) DeleteUpstream(id int32) (*dtos.Upstream, error) {
	model, err := ts.Repository.DeleteUpstream(context.Background(), id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			slog.Error("Failed to get upstream", "error", err)
			return nil, upstreamInternalError
		}
	}

	dto := dtos.Upstream{
		ID:   model.ID,
		Name: model.Name,
	}

	return &dto, nil
}

func (ts UpstreamsService) GetUpstreams(search *dtos.UpstreamSearch) (*dtos.PaginatedUpstreamReponse, error) {
	models, err := ts.Repository.ListUpstreams(context.Background(), db.ListUpstreamsParams{
		Limit:  int32(search.PageSize),
		Offset: int32((search.Page - 1) * search.PageSize),
		Name: "%"+search.Name+"%",
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			slog.Error("Failed to get upstream", "error", err)
			return nil, upstreamInternalError
		}
	}

	length, err := ts.Repository.CountUpstreams(context.Background(), "%"+search.Name+"%")

	if err != nil {
		slog.Error("Failed to get upstream", "error", err)
		return nil, upstreamInternalError
	}

	formated_dtos := []dtos.Upstream{}

	for _, model := range models {
		formated_dtos = append(formated_dtos, dtos.Upstream{
			Name: model.Name,
			ID:   model.ID,
		})
	}

	dto := dtos.PaginatedUpstreamReponse{
		Rows:  formated_dtos,
		Count: length,
	}

	return &dto, nil
}

func (ts UpstreamsService) GetUpstreamUsers(id int32, search *dtos.UserSearch) (*dtos.PaginatedUserReponse, error) {
	models, err := ts.Repository.GetUpstreamUsers(context.Background(), db.GetUpstreamUsersParams{
		Limit:  int32(search.PageSize),
		Offset: int32((search.Page - 1) * search.PageSize),
		UpstreamID: id,
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		} else {
			slog.Error("Failed to get upstream users", "error", err)
			return nil, upstreamInternalError
		}
	}

	length, err := ts.Repository.CountUpstreamUsers(context.Background(), id)

	if err != nil {
		slog.Error("Failed to get upstream users", "error", err)
		return nil, upstreamInternalError
	}

	formated_dtos := []dtos.User{}

	for _, model := range models {
		formated_dtos = append(formated_dtos, dtos.User{
			ID: model.ID,
			Name: model.Name,
			Username: model.Username,
			Email: model.Email,
		})
	}

	dto := dtos.PaginatedUserReponse{
		Rows:  formated_dtos,
		Count: length,
	}

	return &dto, nil
}

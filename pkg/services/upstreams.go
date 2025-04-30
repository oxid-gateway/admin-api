package services

import (
	"context"
	"errors"
	"log/slog"
	"oxid-gateway-admin-api/pkg/db"
	"oxid-gateway-admin-api/pkg/dtos"
	"sync"

	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5"
)

var unique_run_mux = sync.Mutex{}

var internalError = fuego.InternalServerError{Title: "Internal Server Error", Detail: "Upstream server side error"}
var testNotFoundError = fuego.NotFoundError{Title: "Not found", Detail: "Upstream not found"}
var missingTestRunConfigError = fuego.BadRequestError{Title: "Missing Upstream Config", Detail: "Missing required parameters for upstream type"}

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
			return nil, internalError
		}
	}

	dto := dtos.Upstream {
		ID: model.ID,
		Name: model.Name,
	}

	return &dto, nil
}

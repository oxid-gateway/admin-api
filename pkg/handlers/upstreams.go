package handlers

import (
	"log/slog"
	"oxid-gateway-admin-api/pkg/dtos"
	"oxid-gateway-admin-api/pkg/services"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
)

type UpstreamsResources struct {
	UpstreamService *services.UpstreamsService
}

func (rs UpstreamsResources) Routes(s *fuego.Server) {
	fuego.Get(s, "/upstreams/{id}", rs.getUpstream,
		option.Tags("Upstream"),
		option.Summary("Get Upstream By ID"),
		option.Security(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
		option.OperationID("getUpstream"),
	)

	fuego.Delete(s, "/upstreams/{id}", rs.deleteUpstream,
		option.Tags("Upstream"),
		option.Summary("Delete Upstream By ID"),
		option.OperationID("deleteUpstream"),
		option.Security(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)

	fuego.Post(s, "/upstreams", rs.postUpstream,
		option.Tags("Upstream"),
		option.Summary("Craete Upstream"),
		option.OperationID("postUpstream"),
		option.Security(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)

	fuego.Get(s, "/upstreams", rs.getUpstreams,
		dtos.OptionPagination,
		option.Tags("Upstream"),
		option.Summary("Get Upstreams"),
		option.OperationID("getUpstreams"),
		option.Security(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)

	fuego.Get(s, "/upstreams/{id}/users", rs.getUpstreamUsers,
		dtos.OptionPagination,
		option.Tags("Upstream"),
		option.Summary("Get Upstream Users"),
		option.OperationID("getUpstreamUsers"),
		option.Security(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)

	fuego.Put(s, "/upstreams/{id}", rs.putUpstream,
		option.Tags("Upstream"),
		option.Summary("Update Upstream"),
		option.OperationID("putUpstream"),
		option.Security(openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		}),
	)
}

func (ur UpstreamsResources) getUpstream(c fuego.ContextNoBody) (*dtos.Upstream, error) {
	user_context, err := GetRequestContext(c)

	if err != nil {
		return nil, err
	}

	slog.Info("dslkfjds", "user", user_context)

	id := c.PathParamInt("id")

	upstream, err := ur.UpstreamService.GetUpstream(int32(id))

	if err != nil {
		return nil, err
	}

	if upstream == nil {
		return nil, fuego.NotFoundError{Title: "Not found", Detail: "Test not found"}
	}

	return upstream, nil
}

func (ur UpstreamsResources) getUpstreams(c fuego.ContextNoBody) (*dtos.PaginatedUpstreamReponse, error) {
	user_context, err := GetRequestContext(c)

	if err != nil {
		return nil, err
	}

	slog.Info("dslkfjds", "user", user_context)

	return ur.UpstreamService.GetUpstreams(&dtos.UpstreamSearch{
		Page:     c.QueryParamInt("page"),
		PageSize: c.QueryParamInt("pageSize"),
		Name:     c.QueryParam("filter"),
	})
}

func (ur UpstreamsResources) getUpstreamUsers(c fuego.ContextNoBody) (*dtos.PaginatedUserReponse, error) {
	user_context, err := GetRequestContext(c)

	if err != nil {
		return nil, err
	}

	slog.Info("dslkfjds", "user", user_context)

	id := c.QueryParamInt("id")

	return ur.UpstreamService.GetUpstreamUsers(int32(id), &dtos.UserSearch{
		Page:     c.QueryParamInt("page"),
		PageSize: c.QueryParamInt("pageSize"),
	})
}

func (ur UpstreamsResources) deleteUpstream(c fuego.ContextNoBody) (*dtos.Upstream, error) {
	user_context, err := GetRequestContext(c)

	if err != nil {
		return nil, err
	}

	slog.Info("dslkfjds", "user", user_context)

	id := c.PathParamInt("id")

	upstream, err := ur.UpstreamService.DeleteUpstream(int32(id))

	if err != nil {
		return nil, err
	}

	if upstream == nil {
		return nil, fuego.NotFoundError{Title: "Not found", Detail: "Test not found"}
	}

	return upstream, nil
}

func (ur UpstreamsResources) postUpstream(c fuego.ContextWithBody[dtos.UpstreamCreate]) (*dtos.Upstream, error) {
	user_context, err := GetRequestContext(c)

	if err != nil {
		return nil, err
	}

	slog.Info("dslkfjds", "user", user_context)

	body, err := c.Body()

	if err != nil {
		return nil, err
	}

	upstream, err := ur.UpstreamService.CreateUpstream(body)

	if err != nil {
		return nil, err
	}

	if upstream == nil {
		return nil, fuego.NotFoundError{Title: "Not found", Detail: "Test not found"}
	}

	return upstream, nil
}

func (ur UpstreamsResources) putUpstream(c fuego.ContextWithBody[dtos.UpstreamUpdate]) (*dtos.Upstream, error) {
	user_context, err := GetRequestContext(c)

	if err != nil {
		return nil, err
	}

	slog.Info("dslkfjds", "user", user_context)

	id := c.PathParamInt("id")

	body, err := c.Body()

	if err != nil {
		return nil, err
	}

	upstream, err := ur.UpstreamService.UpdateUpstream(int32(id), body)

	if err != nil {
		return nil, err
	}

	if upstream == nil {
		return nil, fuego.NotFoundError{Title: "Not found", Detail: "Test not found"}
	}

	return upstream, nil
}

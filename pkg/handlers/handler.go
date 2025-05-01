package handlers

import (
	"oxid-gateway-admin-api/pkg/dtos"
	"oxid-gateway-admin-api/pkg/services"
	"strconv"

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
		option.OperationID("getUpstream"),
	)

	fuego.Post(s, "/upstreams", rs.postUpstream,
		option.Tags("Upstream"),
		option.Summary("Craete Upstream"),
		option.OperationID("postUpstream"),
	)

	fuego.Put(s, "/upstreams/{id}", rs.putUpstream,
		option.Tags("Upstream"),
		option.Summary("Update Upstream"),
		option.OperationID("putUpstream"),
	)
}

func (ur UpstreamsResources) getUpstream(c fuego.ContextNoBody) (*dtos.Upstream, error) {
	id, err := strconv.ParseInt(c.PathParam("id"), 10, 64)

	if err != nil {
		return nil, err
	}

	upstream, err := ur.UpstreamService.GetUpstream(int32(id))

	if err != nil {
		return nil, err
	}

	if upstream == nil {
		return nil, fuego.NotFoundError{Title: "Not found", Detail: "Test not found"}
	}

	return upstream, nil
}

func (ur UpstreamsResources) postUpstream(c fuego.ContextWithBody[dtos.UpstreamCreate]) (*dtos.Upstream, error) {
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
	id, err := strconv.ParseInt(c.PathParam("id"), 10, 64)

	if err != nil {
		return nil, err
	}

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

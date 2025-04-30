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
}

func (ur UpstreamsResources) getUpstream(c fuego.ContextNoBody) (*dtos.Upstream, error) {
	id, err := strconv.ParseInt(c.PathParam("id"), 10 , 64)

	if err != nil {
		panic(err)
	}

	upstream, err := ur.UpstreamService.GetUpstream(int32(id))

	if err != nil {
		panic(err)
	}


	if upstream == nil {
		return nil, fuego.NotFoundError{Title: "Not found", Detail: "Test not found"}
	}

	return upstream, nil
}

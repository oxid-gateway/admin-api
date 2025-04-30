package dtos

import (
	"github.com/go-fuego/fuego/option"
	"github.com/go-fuego/fuego/param"
)

type PaginatedResponse[Res any] struct {
	Rows  []Res `json:"rows" validate:"required"`
	Count int64 `json:"count" validate:"required"`
}

var OptionPagination = option.Group(
	option.QueryInt("page", "Page number", param.Default(1)),
	option.QueryInt("pageSize", "Number of items per page", param.Default(10), param.Example("10 items per page", 10)),
)

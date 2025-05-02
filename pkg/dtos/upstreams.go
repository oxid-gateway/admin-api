package dtos

type Upstream struct {
	ID    int32 `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
}

type UpstreamCreate struct {
	Name    string `json:"name" validate:"required"`
}

type UpstreamUpdate struct {
	Name string `json:"name"`
}

type UpstreamSearch struct {
	Page     int
	PageSize int
	Name     string
}


type PaginatedUpstreamReponse PaginatedResponse[Upstream]

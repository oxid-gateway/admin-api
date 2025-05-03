package dtos

type User struct {
	Name  string `json:"name" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Email  string `json:"email" validate:"required"`
}

type UserSearch struct {
	Page     int
	PageSize int
}

type PaginatedUserReponse PaginatedResponse[User]

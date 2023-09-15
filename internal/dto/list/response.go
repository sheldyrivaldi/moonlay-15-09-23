package listdto

import (
	"moonlay-todolist/internal/abstraction"
	"moonlay-todolist/internal/model"
)

type FindAllResponse struct {
	Data           []model.List               `json:"data"`
	PaginationInfo abstraction.PaginationInfo `json:"pagination_info"`
	SubList        bool                       `json:"-"`
}

type FindByIDResponse struct {
	Data model.List `json:"data"`
}

type CreateResponse struct {
	Data model.List `json:"data"`
}

type UpdateByIDResponse struct {
	Data model.List `json:"data"`
}

type DeleteByIDResponse struct {
	Data model.List `json:"data"`
}

package sublistdto

import (
	"moonlay-todolist/internal/abstraction"
	"moonlay-todolist/internal/model"
)

type FindAllResponse struct {
	Data           []model.SubList            `json:"data"`
	PaginationInfo abstraction.PaginationInfo `json:"pagination_info"`
}

type FindByIDResponse struct {
	Data model.SubList `json:"data"`
}

type CreateResponse struct {
	Data model.SubList `json:"data"`
}

type UpdateByIDResponse struct {
	Data model.SubList `json:"data"`
}

type DeleteByIDResponse struct {
	Data model.SubList `json:"data"`
}

package sublistdto

import (
	"moonlay-todolist/internal/abstraction"
	"moonlay-todolist/internal/model"
)

type FindAllRequest struct {
	abstraction.Pagination
	ListID      string `validate:"required"`
	Title       string
	Description string
}

type FindByIDRequest struct {
	ListID string `param:"listID" validate:"required"`
	ID     string `param:"id" validate:"required"`
}

type CreateRequest struct {
	ListID      string `json:"list_id" validate:"required"`
	Title       string `json:"title" validate:"max=100"`
	Description string `json:"description" validate:"max=1000"`
}

type UpdateByIDRequest struct {
	ID     string `param:"id" validate:"required"`
	ListID string `param:"listID" validate:"required"`
	model.SubList
}

type DeleteByIDRequest struct {
	ID     string `param:"id" validate:"required"`
	ListID string `param:"listID" validate:"required"`
}

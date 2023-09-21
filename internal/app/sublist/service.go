package sublist

import (
	"errors"
	sublistdto "moonlay-todolist/internal/dto/sublist"
	"moonlay-todolist/internal/factory"
	"moonlay-todolist/internal/model"
	"moonlay-todolist/internal/repository"

	"github.com/google/uuid"
)

type IService interface {
	FindAll(payload *sublistdto.FindAllRequest) (*sublistdto.FindAllResponse, error)
	FindByID(payload *sublistdto.FindByIDRequest) (*sublistdto.FindByIDResponse, error)
	Create(payload *sublistdto.CreateRequest) (*sublistdto.CreateResponse, error)
	UpdateByID(payload *sublistdto.UpdateByIDRequest, files []string) (*sublistdto.UpdateByIDResponse, error)
	DeleteByID(payload *sublistdto.DeleteByIDRequest) (*sublistdto.DeleteByIDResponse, error)
}

type service struct {
	Repository repository.ISubListRepository
}

func NewSubListService(f *factory.Factory) *service {
	repository := f.SubListRepository
	return &service{repository}
}

func (l service) FindAll(payload *sublistdto.FindAllRequest) (*sublistdto.FindAllResponse, error) {
	var result *sublistdto.FindAllResponse
	var data *[]model.SubList

	data, info, err := l.Repository.FindAll(&payload.ListID, &payload.Description, &payload.Title, &payload.Pagination)
	if err != nil {
		return result, errors.New("failed to find data")
	}

	result = &sublistdto.FindAllResponse{
		Data:           *data,
		PaginationInfo: *info,
	}

	return result, nil
}

func (l service) FindByID(payload *sublistdto.FindByIDRequest) (*sublistdto.FindByIDResponse, error) {
	var result *sublistdto.FindByIDResponse

	data, err := l.Repository.FindByID(payload.ListID, payload.ID)
	if err != nil {
		return result, err
	}

	result = &sublistdto.FindByIDResponse{
		Data: *data,
	}

	return result, nil
}

func (l service) Create(payload *sublistdto.CreateRequest, files []string) (*sublistdto.CreateResponse, error) {
	var result *sublistdto.CreateResponse
	list := &model.SubList{}
	ID, err := uuid.NewRandom()
	if err != nil {
		return result, err
	}

	list.ID = ID.String()
	list.ListID = payload.ListID
	list.Title = payload.Title
	list.Description = payload.Description

	data, err := l.Repository.Create(list)
	if err != nil {
		return result, err
	}

	for _, v := range files {
		_, err := l.Repository.CreateFile(v, list.ID)
		if err != nil {
			return result, err
		}

	}

	result = &sublistdto.CreateResponse{
		Data: *data,
	}

	return result, nil
}

func (l service) UpdateByID(payload *sublistdto.UpdateByIDRequest, files []string) (*sublistdto.UpdateByIDResponse, error) {
	var result *sublistdto.UpdateByIDResponse
	list := &model.SubList{}

	if payload.Title != "" {
		list.Title = payload.Title
	}
	if payload.Description != "" {
		list.Description = payload.Title
	}

	data, err := l.Repository.UpdateByID(payload.ListID, payload.ID, list)
	if err != nil {
		return result, err
	}

	err = l.Repository.DeleteFileBySubListID(payload.ID)
	if err != nil {
		return result, err
	}

	for _, v := range files {
		_, err := l.Repository.CreateFile(v, payload.ID)
		if err != nil {
			return result, err
		}

	}

	result = &sublistdto.UpdateByIDResponse{
		Data: *data,
	}

	return result, nil
}

func (l service) DeleteByID(payload *sublistdto.DeleteByIDRequest) (*sublistdto.DeleteByIDResponse, error) {
	var result *sublistdto.DeleteByIDResponse

	errDel := l.Repository.DeleteFileBySubListID(payload.ID)
	if errDel != nil {
		return result, errDel
	}

	data, err := l.Repository.DeleteByID(payload.ListID, payload.ID)
	if err != nil {
		return result, err
	}

	result = &sublistdto.DeleteByIDResponse{
		Data: *data,
	}

	return result, nil
}

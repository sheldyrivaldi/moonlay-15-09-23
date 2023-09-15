package list

import (
	"errors"
	listdto "moonlay-todolist/internal/dto/list"
	"moonlay-todolist/internal/factory"
	"moonlay-todolist/internal/model"
	"moonlay-todolist/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IService interface {
	FindAll(payload *listdto.FindAllRequest) (*listdto.FindAllResponse, error)
	FindByID(payload *listdto.FindByIDRequest) (*listdto.FindByIDResponse, error)
	Create(payload *listdto.CreateRequest) (*listdto.CreateResponse, error)
	UpdateByID(payload *listdto.UpdateByIDRequest, files []string) (*listdto.UpdateByIDResponse, error)
	DeleteByID(payload *listdto.DeleteByIDRequest) (*listdto.DeleteByIDResponse, error)
}

type service struct {
	Repository repository.IListRepository
	Db         *gorm.DB
}

func NewListService(f *factory.Factory) *service {
	repository := f.ListRepository
	db := f.Db
	return &service{repository, db}
}

func (l service) FindAll(payload *listdto.FindAllRequest) (*listdto.FindAllResponse, error) {
	var result *listdto.FindAllResponse
	var data *[]model.List

	data, info, err := l.Repository.FindAll(&payload.Sublist, &payload.Description, &payload.Title, &payload.Pagination)
	if err != nil {
		return result, errors.New("failed to find data")
	}

	result = &listdto.FindAllResponse{
		Data:           *data,
		PaginationInfo: *info,
		SubList:        payload.Sublist,
	}

	return result, nil
}

func (l service) FindByID(payload *listdto.FindByIDRequest) (*listdto.FindByIDResponse, error) {
	var result *listdto.FindByIDResponse

	data, err := l.Repository.FindByID(payload.ID)
	if err != nil {
		return result, err
	}

	result = &listdto.FindByIDResponse{
		Data: *data,
	}

	return result, nil
}

func (l service) Create(payload *listdto.CreateRequest, files []string) (*listdto.CreateResponse, error) {
	var result *listdto.CreateResponse
	list := &model.List{}
	ID, err := uuid.NewRandom()
	if err != nil {
		return result, err
	}

	list.ID = ID.String()
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

	result = &listdto.CreateResponse{
		Data: *data,
	}

	return result, nil
}

func (l service) UpdateByID(payload *listdto.UpdateByIDRequest, files []string) (*listdto.UpdateByIDResponse, error) {
	var result *listdto.UpdateByIDResponse
	list := &model.List{}
	if payload.Title != "" {
		list.Title = payload.Title
	}
	if payload.Description != "" {
		list.Description = payload.Title
	}

	data, err := l.Repository.UpdateByID(payload.ID, list)
	if err != nil {
		return result, err
	}

	err = l.Repository.DeleteFileByListID(payload.ID)
	if err != nil {
		return result, err
	}

	for _, v := range files {
		_, err := l.Repository.CreateFile(v, payload.ID)
		if err != nil {
			return result, err
		}

	}

	result = &listdto.UpdateByIDResponse{
		Data: *data,
	}

	return result, nil
}

func (l service) DeleteByID(payload *listdto.DeleteByIDRequest) (*listdto.DeleteByIDResponse, error) {
	var result *listdto.DeleteByIDResponse

	errDel := l.Repository.DeleteFileByListID(payload.ID)
	if errDel != nil {
		return result, errDel
	}

	data, err := l.Repository.DeleteByID(payload.ID)
	if err != nil {
		return result, err
	}

	result = &listdto.DeleteByIDResponse{
		Data: *data,
	}

	return result, nil
}

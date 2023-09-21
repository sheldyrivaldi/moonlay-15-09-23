package sublist

import (
	"errors"
	"moonlay-todolist/internal/abstraction"
	sublistdto "moonlay-todolist/internal/dto/sublist"
	"moonlay-todolist/internal/factory"
	"moonlay-todolist/internal/mocks"
	"moonlay-todolist/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindAll(t *testing.T) {
	mockRepo := new(mocks.ISubListRepository)
	service := NewSubListService(&factory.Factory{SubListRepository: mockRepo})

	t.Run("Success", func(t *testing.T) {
		expectedResponse := &sublistdto.FindAllResponse{
			Data: []model.SubList{
				{
					Title:       "opopop",
					Description: "ssfess",
				},
			},
			PaginationInfo: abstraction.PaginationInfo{},
		}

		mockRepo.On("FindAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&[]model.SubList{expectedResponse.Data[0]}, &expectedResponse.PaginationInfo, nil).Once()

		result, err := service.FindAll(&sublistdto.FindAllRequest{})

		assert.NoError(t, err)

		assert.Equal(t, result, expectedResponse)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockRepo.On("FindAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil, errors.New("error")).Once()

		result, err := service.FindAll(&sublistdto.FindAllRequest{})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

}

func TestFindByID(t *testing.T) {
	mockRepo := new(mocks.ISubListRepository)
	service := NewSubListService(&factory.Factory{SubListRepository: mockRepo})

	expectedResponse := &sublistdto.FindByIDResponse{
		Data: model.SubList{
			ListID:      "safaf",
			Title:       "opopop",
			Description: "ssfess",
		},
	}

	t.Run("Success", func(t *testing.T) {

		mockRepo.On("FindByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()

		result, err := service.FindByID(&sublistdto.FindByIDRequest{
			ID: "safaf",
		})

		assert.NoError(t, err)
		assert.Equal(t, result, expectedResponse)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {

		mockRepo.On("FindByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("failed to find data")).Once()

		result, err := service.FindByID(&sublistdto.FindByIDRequest{
			ID: "safaf",
		})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockRepo := new(mocks.ISubListRepository)
	service := NewSubListService(&factory.Factory{SubListRepository: mockRepo})

	expectedResponse := &sublistdto.CreateResponse{
		Data: model.SubList{
			ListID:      "safaf",
			Title:       "opopop",
			Description: "ssfess",
		},
	}
	expectedResponseFile := model.SubListFile{
		Link: "Link 1",
	}

	t.Run("Success", func(t *testing.T) {

		mockRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()
		mockRepo.On("CreateFile", mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponseFile, nil).Once()

		result, err := service.Create(&sublistdto.CreateRequest{
			ListID:      "safaf",
			Title:       "opopop",
			Description: "ssfess",
		}, []string{"Link 1"})

		assert.NoError(t, err)
		assert.Equal(t, result, expectedResponse)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Create File", func(t *testing.T) {

		mockRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()
		mockRepo.On("CreateFile", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()

		result, err := service.Create(&sublistdto.CreateRequest{
			ListID:      "safaf",
			Title:       "opopop",
			Description: "ssfess",
		}, []string{"Link 1"})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Create", func(t *testing.T) {

		mockRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(&model.SubList{}, errors.New("error")).Once()

		result, err := service.Create(&sublistdto.CreateRequest{
			ListID:      "safaf",
			Title:       "opopop",
			Description: "ssfess",
		}, []string{"Link 1"})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateyByID(t *testing.T) {
	mockRepo := new(mocks.ISubListRepository)
	service := NewSubListService(&factory.Factory{SubListRepository: mockRepo})

	expectedResponse := &sublistdto.UpdateByIDResponse{
		Data: model.SubList{
			ListID:      "safaf",
			Title:       "opopop",
			Description: "ssfess",
		},
	}

	expectedResponseFile := model.SubListFile{
		Link: "Link 1",
	}

	t.Run("Success", func(t *testing.T) {

		mockRepo.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()
		mockRepo.On("DeleteFileBySubListID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		mockRepo.On("CreateFile", mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponseFile, nil).Once()

		result, err := service.UpdateByID(&sublistdto.UpdateByIDRequest{
			ID: "safasf",
			SubList: model.SubList{
				Title:       "opopop",
				Description: "ssfess",
			},
		}, []string{"Link 1"})

		assert.NoError(t, err)
		assert.Equal(t, result, expectedResponse)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Update", func(t *testing.T) {

		mockRepo.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()

		result, err := service.UpdateByID(&sublistdto.UpdateByIDRequest{
			ID: "safasf",
			SubList: model.SubList{
				Title:       "opopop",
				Description: "ssfess",
			},
		}, []string{"Link 1"})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Delete File", func(t *testing.T) {

		mockRepo.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()
		mockRepo.On("DeleteFileBySubListID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		mockRepo.On("CreateFile", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()

		result, err := service.UpdateByID(&sublistdto.UpdateByIDRequest{
			ID: "safasf",
			SubList: model.SubList{
				Title:       "opopop",
				Description: "ssfess",
			},
		}, []string{"Link 1"})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Create File", func(t *testing.T) {

		mockRepo.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()
		mockRepo.On("DeleteFileBySubListID", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		result, err := service.UpdateByID(&sublistdto.UpdateByIDRequest{
			ID: "safasf",
			SubList: model.SubList{
				Title:       "opopop",
				Description: "ssfess",
			},
		}, []string{"Link 1"})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteByID(t *testing.T) {
	mockRepo := new(mocks.ISubListRepository)
	service := NewSubListService(&factory.Factory{SubListRepository: mockRepo})

	expectedResponse := &sublistdto.DeleteByIDResponse{
		Data: model.SubList{
			ListID:      "safaf",
			Title:       "opopop",
			Description: "ssfess",
		},
	}

	t.Run("Success", func(t *testing.T) {

		mockRepo.On("DeleteFileBySubListID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		mockRepo.On("DeleteByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()

		result, err := service.DeleteByID(&sublistdto.DeleteByIDRequest{
			ID: "asfasf",
		})

		assert.NoError(t, err)
		assert.Equal(t, result, expectedResponse)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Delete File", func(t *testing.T) {

		mockRepo.On("DeleteFileBySubListID", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		result, err := service.DeleteByID(&sublistdto.DeleteByIDRequest{
			ID: "asfasf",
		})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Delete By ID", func(t *testing.T) {

		mockRepo.On("DeleteFileBySubListID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		mockRepo.On("DeleteByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()

		result, err := service.DeleteByID(&sublistdto.DeleteByIDRequest{
			ID: "asfasf",
		})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

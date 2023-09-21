package list

import (
	"errors"
	"moonlay-todolist/internal/abstraction"
	listdto "moonlay-todolist/internal/dto/list"
	"moonlay-todolist/internal/factory"
	"moonlay-todolist/internal/mocks"
	"moonlay-todolist/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindAll(t *testing.T) {
	mockRepo := new(mocks.IListRepository)
	service := NewListService(&factory.Factory{ListRepository: mockRepo})

	t.Run("Success", func(t *testing.T) {
		expectedResponse := &listdto.FindAllResponse{
			Data: []model.List{
				{
					Title:       "opopop",
					Description: "ssfess",
				},
			},
			PaginationInfo: abstraction.PaginationInfo{},
		}

		mockRepo.On("FindAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&[]model.List{expectedResponse.Data[0]}, &expectedResponse.PaginationInfo, nil).Once()

		result, err := service.FindAll(&listdto.FindAllRequest{})

		assert.NoError(t, err)

		assert.Equal(t, result, expectedResponse)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {

		mockRepo.On("FindAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil, errors.New("error")).Once()

		result, err := service.FindAll(&listdto.FindAllRequest{})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

}

func TestFindByID(t *testing.T) {
	mockRepo := new(mocks.IListRepository)
	service := NewListService(&factory.Factory{ListRepository: mockRepo})

	expectedResponse := &listdto.FindByIDResponse{
		Data: model.List{
			Title:       "opopop",
			Description: "ssfess",
		},
	}

	t.Run("Success", func(t *testing.T) {

		mockRepo.On("FindByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()

		result, err := service.FindByID(&listdto.FindByIDRequest{
			ID: "safaf",
		})

		assert.NoError(t, err)
		assert.Equal(t, result, expectedResponse)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {

		mockRepo.On("FindByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("failed to find data")).Once()

		result, err := service.FindByID(&listdto.FindByIDRequest{
			ID: "safaf",
		})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	mockRepo := new(mocks.IListRepository)
	service := NewListService(&factory.Factory{ListRepository: mockRepo})

	expectedResponse := &listdto.CreateResponse{
		Data: model.List{
			Title:       "opopop",
			Description: "ssfess",
		},
	}
	expectedResponseFile := model.ListFile{
		Link: "Link 1",
	}

	t.Run("Success", func(t *testing.T) {

		mockRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()
		mockRepo.On("CreateFile", mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponseFile, nil).Once()

		result, err := service.Create(&listdto.CreateRequest{
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

		result, err := service.Create(&listdto.CreateRequest{
			Title:       "opopop",
			Description: "ssfess",
		}, []string{"Link 1"})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Create", func(t *testing.T) {

		mockRepo.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(&model.List{}, errors.New("error")).Once()

		result, err := service.Create(&listdto.CreateRequest{
			Title:       "opopop",
			Description: "ssfess",
		}, []string{"Link 1"})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateyByID(t *testing.T) {
	mockRepo := new(mocks.IListRepository)
	service := NewListService(&factory.Factory{ListRepository: mockRepo})

	expectedResponse := &listdto.UpdateByIDResponse{
		Data: model.List{
			Title:       "opopop",
			Description: "ssfess",
		},
	}

	expectedResponseFile := model.ListFile{
		Link: "Link 1",
	}

	t.Run("Success", func(t *testing.T) {

		mockRepo.On("UpdateByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()
		mockRepo.On("DeleteFileByListID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		mockRepo.On("CreateFile", mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponseFile, nil).Once()

		result, err := service.UpdateByID(&listdto.UpdateByIDRequest{
			ID: "safasf",
			List: model.List{
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

		result, err := service.UpdateByID(&listdto.UpdateByIDRequest{
			ID: "safasf",
			List: model.List{
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
		mockRepo.On("DeleteFileByListID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		mockRepo.On("CreateFile", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()

		result, err := service.UpdateByID(&listdto.UpdateByIDRequest{
			ID: "safasf",
			List: model.List{
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
		mockRepo.On("DeleteFileByListID", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		result, err := service.UpdateByID(&listdto.UpdateByIDRequest{
			ID: "safasf",
			List: model.List{
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
	mockRepo := new(mocks.IListRepository)
	service := NewListService(&factory.Factory{ListRepository: mockRepo})

	expectedResponse := &listdto.DeleteByIDResponse{
		Data: model.List{
			Title:       "opopop",
			Description: "ssfess",
		},
	}

	t.Run("Success", func(t *testing.T) {

		mockRepo.On("DeleteFileByListID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		mockRepo.On("DeleteByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()

		result, err := service.DeleteByID(&listdto.DeleteByIDRequest{
			ID: "asfasf",
		})

		assert.NoError(t, err)
		assert.Equal(t, result, expectedResponse)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Delete File", func(t *testing.T) {

		mockRepo.On("DeleteFileByListID", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		result, err := service.DeleteByID(&listdto.DeleteByIDRequest{
			ID: "asfasf",
		})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Error Delete By ID", func(t *testing.T) {

		mockRepo.On("DeleteFileByListID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		mockRepo.On("DeleteByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()

		result, err := service.DeleteByID(&listdto.DeleteByIDRequest{
			ID: "asfasf",
		})

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

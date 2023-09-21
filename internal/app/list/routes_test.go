package list

import (
	"moonlay-todolist/internal/abstraction"
	listdto "moonlay-todolist/internal/dto/list"
	"moonlay-todolist/internal/factory"
	"moonlay-todolist/internal/mocks"
	"moonlay-todolist/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRoute(t *testing.T) {
	e := echo.New()

	mockService := new(mocks.IListRepository)
	handler := NewHandler(&factory.Factory{ListRepository: mockService})

	g := e.Group("/list")
	handler.Route(g)

	testCases := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "FindAll",
			method:         http.MethodGet,
			path:           "/list",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "FindByID",
			method:         http.MethodGet,
			path:           "/list/listID",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Create",
			method:         http.MethodPost,
			path:           "/list",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "UpdateByID",
			method:         http.MethodPut,
			path:           "/list/listID",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "DeleteByID",
			method:         http.MethodDelete,
			path:           "/list/listID",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			rec := httptest.NewRecorder()

			expectedResponse1 := &listdto.FindAllResponse{
				Data: []model.List{
					{
						Title:       "opopop",
						Description: "ssfess",
					},
				},
				PaginationInfo: abstraction.PaginationInfo{},
			}

			mockService.On("FindAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&[]model.List{expectedResponse1.Data[0]}, &expectedResponse1.PaginationInfo, nil).Once()

			expectedResponse := &listdto.FindByIDResponse{
				Data: model.List{
					Title:       "opopop",
					Description: "ssfess",
				},
			}

			mockService.On("FindByID", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&expectedResponse.Data, nil).Once()

			e.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
		})
	}
}

package sublist

import (
	"moonlay-todolist/internal/abstraction"
	sublistdto "moonlay-todolist/internal/dto/sublist"
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

	mockService := new(mocks.ISubListRepository)
	handler := NewHandler(&factory.Factory{SubListRepository: mockService})

	g := e.Group("/list/listID/sub")
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
			path:           "/list/listID/sub",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "FindByID",
			method:         http.MethodGet,
			path:           "/list/listID/sub/sublistID",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Create",
			method:         http.MethodPost,
			path:           "/list/listID/sub",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "UpdateByID",
			method:         http.MethodPut,
			path:           "/list/listID/sub/sublistID",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "DeleteByID",
			method:         http.MethodDelete,
			path:           "/list/listID/sub/sublistID",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			rec := httptest.NewRecorder()

			expectedResponse1 := &sublistdto.FindAllResponse{
				Data: []model.SubList{
					{
						Title:       "opopop",
						Description: "ssfess",
					},
				},
				PaginationInfo: abstraction.PaginationInfo{},
			}

			mockService.On("FindAll", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&[]model.SubList{expectedResponse1.Data[0]}, &expectedResponse1.PaginationInfo, nil).Once()

			expectedResponse := &sublistdto.FindByIDResponse{
				Data: model.SubList{
					ListID:      "safaf",
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

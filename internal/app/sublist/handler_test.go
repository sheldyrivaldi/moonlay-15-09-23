package sublist

import (
	"io"
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

func getMockedContext(method, target string, body io.Reader) echo.Context {
	e := echo.New()
	req := httptest.NewRequest(method, target, body)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c
}

func TestFindAllHandler(t *testing.T) {
	UrlExample := []string{
		"/list/8d037741-8335-4fd6-bc4a-e73c6c466a74/sub",
		"/list/8d037741-8335-4fd6-bc4a-e73c6c466a74/sub?page=1",
		"/list/8d037741-8335-4fd6-bc4a-e73c6c466a74/sub?pageSize=5",
		"/list/8d037741-8335-4fd6-bc4a-e73c6c466a74/sub?sortBy=created_at",
		"/list/8d037741-8335-4fd6-bc4a-e73c6c466a74/sub?sort=asc",
		"/list/8d037741-8335-4fd6-bc4a-e73c6c466a74/sub?title=test",
		"/list/8d037741-8335-4fd6-bc4a-e73c6c466a74/sub?description=testing",
	}

	for _, v := range UrlExample {
		t.Run("NoError", func(t *testing.T) {
			mockService := new(mocks.ISubListService)

			handler := NewHandler(&factory.Factory{})

			expectedResponse := &sublistdto.FindAllResponse{
				Data: []model.SubList{
					{
						Title:       "title",
						Description: "description",
					},
				},
			}

			ctx := getMockedContext(http.MethodGet, v, nil)

			mockService.On("FindAll", mock.Anything).Return(expectedResponse, nil)

			err := handler.FindAll(ctx)

			assert.NoError(t, err)
		})
	}
}

func TestFindByIDHandler(t *testing.T) {

	handler := NewHandler(&factory.Factory{})

	ctx := getMockedContext(http.MethodGet, "/list/aaaa/sub/aaa", nil)

	err := handler.FindByID(ctx)

	assert.NoError(t, err)

}

func TestCreateHandler(t *testing.T) {
	handler := NewHandler(&factory.Factory{})

	ctx := getMockedContext(http.MethodPost, "/list/listID/sub", nil)

	err := handler.Create(ctx)

	assert.NoError(t, err)

}

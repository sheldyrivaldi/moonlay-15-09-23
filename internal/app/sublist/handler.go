package sublist

import (
	"fmt"
	resultdto "moonlay-todolist/internal/dto/result"
	sublistdto "moonlay-todolist/internal/dto/sublist"
	"moonlay-todolist/internal/factory"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IHandler interface {
	FindAll(c echo.Context) error
	FindByID(c echo.Context) error
	Create(c echo.Context) error
	UpdateByID(c echo.Context) error
	DeleteByID(c echo.Context) error
}

type handler struct {
	service *service
}

func NewHandler(f *factory.Factory) *handler {
	service := NewSubListService(f)
	return &handler{service}
}

func (h *handler) FindAll(c echo.Context) error {
	payload := new(sublistdto.FindAllRequest)

	if c.QueryParam("page") != "" {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		payload.Page = &page
	}

	if c.QueryParam("pageSize") != "" {
		pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
		payload.PageSize = &pageSize
	}

	if c.QueryParam("sortBy") != "" {
		sortBy := c.QueryParam("sortBy")
		payload.SortBy = &sortBy
	}

	if c.QueryParam("sort") != "" {
		sort := c.QueryParam("sort")
		payload.Sort = &sort
	}

	if c.QueryParam("title") != "" {
		payload.Title = c.QueryParam("title")
	}

	if c.QueryParam("description") != "" {
		payload.Description = c.QueryParam("description")
	}

	if c.Param("listID") != "" {
		payload.ListID = c.Param("listID")
	}

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Failed to binding input data!",
		})
	}

	err2 := c.Validate(payload)
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Failed to validate input data!",
		})
	}

	result, err := h.service.FindAll(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to find data at service!",
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	})
}

func (h handler) FindByID(c echo.Context) error {
	payload := new(sublistdto.FindByIDRequest)

	ID := c.Param("id")
	ListID := c.Param("listID")
	payload.ID = ID
	payload.ListID = ListID

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Failed to binding input data!",
		})
	}

	err = c.Validate(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Failed to validate input data!",
		})
	}

	result, err := h.service.FindByID(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to find data at service!",
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	})
}

func (h handler) Create(c echo.Context) error {
	payload := new(sublistdto.CreateRequest)

	ListID := c.Param("listID")
	title := c.FormValue("title")
	description := c.FormValue("description")

	payload.ListID = ListID
	payload.Title = title
	payload.Description = description

	err := c.Validate(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Failed to validate input data!",
		})
	}

	dataFile := c.Get("dataFiles").([]string)
	result, err := h.service.Create(payload, dataFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to create data!",
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	})
}

func (h handler) UpdateByID(c echo.Context) error {
	payload := new(sublistdto.UpdateByIDRequest)
	payloadData := new(sublistdto.FindByIDRequest)

	ID := c.Param("id")
	listID := c.Param("listID")

	payload.ID = ID
	payload.ListID = listID
	payloadData.ID = ID
	payloadData.ListID = listID

	data, err := h.service.FindByID(payloadData)
	if data == nil {
		return c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: fmt.Sprintf("No data with ID %s !", payloadData.ID),
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to get data!",
		})
	}

	title := c.FormValue("title")
	description := c.FormValue("description")

	payload.Title = title
	payload.Description = description

	err = c.Validate(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Failed to validate input data!",
		})
	}

	dataFile := c.Get("dataFiles").([]string)
	result, err := h.service.UpdateByID(payload, dataFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to update data!",
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	})
}

func (h handler) DeleteByID(c echo.Context) error {
	payload := new(sublistdto.DeleteByIDRequest)
	payloadData := new(sublistdto.FindByIDRequest)

	ID := c.Param("id")
	listID := c.Param("listID")

	payload.ID = ID
	payloadData.ID = ID
	payload.ListID = listID
	payloadData.ListID = listID

	data, err := h.service.FindByID(payloadData)
	if data == nil {
		return c.JSON(http.StatusNotFound, resultdto.ErrorResult{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: fmt.Sprintf("No data with ID %s !", payloadData.ID),
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to get data!",
		})
	}

	err = c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Failed to binding input data!",
		})
	}

	err = c.Validate(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Failed to validate input data!",
		})
	}

	result, err := h.service.DeleteByID(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Status:  "Internal Server Error",
			Message: "Failed to delete data!",
		})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	})
}

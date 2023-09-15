package list

import (
	echoCloudinary "moonlay-todolist/internal/pkg/cloudinary"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(e *echo.Group) {
	e.GET("", h.FindAll)
	e.GET("/:id", h.FindByID)
	e.POST("", echoCloudinary.UploadFile(h.Create))
	e.PUT("/:id", echoCloudinary.UploadFile(h.UpdateByID))
	e.DELETE("/:id", h.DeleteByID)
}

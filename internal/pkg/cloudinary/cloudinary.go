package echoCloudinary

import (
	"context"
	resultdto "moonlay-todolist/internal/dto/result"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var ctx = context.Background()
		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")
		cld, err := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
				Code:    http.StatusInternalServerError,
				Status:  "Internal Server Error",
				Message: "Failed connect to cloudinary",
			})
		}

		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
				Code:    http.StatusBadRequest,
				Status:  "Bad Request",
				Message: "Failed read file!",
			})
		}

		files := form.File["files"]
		allowedExtensions := map[string]bool{
			".txt": true,
			".pdf": true,
		}

		var dataFiles []string
		for _, file := range files {
			ext := filepath.Ext(file.Filename)
			if !allowedExtensions[ext] {
				return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
					Code:    http.StatusBadRequest,
					Status:  "Bad Request",
					Message: "The file extension is wrong. Allowed file extensions are .txt and .pdf",
				})
			}

			src, err := file.Open()
			if err != nil {
				return c.JSON(http.StatusBadRequest, err)
			}
			defer src.Close()

			resp, errUpload := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "moonlay-test"})
			if errUpload != nil {
				return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
					Code:    http.StatusInternalServerError,
					Status:  "Internal Server Error",
					Message: "Filed upload to file",
				})
			}

			dataFiles = append(dataFiles, resp.SecureURL)
		}
		c.Set("dataFiles", dataFiles)
		return next(c)
	}
}

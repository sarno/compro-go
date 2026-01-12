package handler

import (
	"bytes"
	"compro/config"
	"compro/internal/adapter/handler/response"
	"compro/internal/adapter/storage"
	"compro/utils/middleware"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type IUploadImageHandler interface {
	UploadImage(c echo.Context) error
}

type UploadImageHandler struct {
	storageService storage.ISupabase
}

// UploadImage implements [IUploadImageHandler].
func (u *UploadImageHandler) UploadImage(c echo.Context) error {
	var (
		respError = response.ErrorResponseDefault{}
		resp      = response.DefaultSuccessResponse{}
	)

	file, err := c.FormFile("file")
	if err != nil {
		log.Errorf("Error getting file: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(400, respError)
	}

	src, err := file.Open()
	if err != nil {
		log.Errorf("Error opening file: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(400, respError)
	}

	defer src.Close()
	fileBuffer := new(bytes.Buffer)
	_, err = io.Copy(fileBuffer, src)

	if err != nil {
		log.Errorf("Error copying file: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(400, respError)
	}

	newFileName := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), getExtension(file.Filename))
	uploadPath := fmt.Sprintf("public/uploads/%s", newFileName)
	url, err := u.storageService.UploadFile(uploadPath, fileBuffer)

	if err != nil {
		log.Errorf("Error uploading file: %v", err)
		respError.Meta.Message = err.Error()
		respError.Meta.Status = false
		return c.JSON(400, respError)
	}

	resp.Meta.Status = true
	resp.Meta.Message = "Success upload image"
	resp.Data = map[string]string{
		"url": url,
	}

	resp.Pagination = nil
	return c.JSON(http.StatusCreated, resp)
}

func NewUploadImageHandler(e *echo.Echo, storageService storage.ISupabase, cfg *config.Config) IUploadImageHandler {
	res := &UploadImageHandler{
		storageService: storageService,
	}

	mid := middleware.NewMiddleware(cfg)
	e.POST("/upload-image", res.UploadImage, mid.CheckToken())

	return res
}

func getExtension(fileName string) string {
	ext := "." + fileName[len(fileName)-3:] // Ambil 3 karakter terakhir untuk ekstensi
	if len(fileName) > 4 && fileName[len(fileName)-4] == '.' {
		ext = "." + fileName[len(fileName)-4:]
	}
	return ext
}

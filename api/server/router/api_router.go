package router

import (
	"github.com/labstack/echo"
	"github.com/takutakukatokatojapan/image_analysis_api/api/server/handler"
)

type (
	Router interface {
		Route(e *echo.Echo)
	}
	V1Router struct {
		uploadFileHandler handler.UploadFileHandler
	}
)

func NewV1Router(uploadFileHandler handler.UploadFileHandler) Router {
	return &V1Router{uploadFileHandler: uploadFileHandler}
}

func (r V1Router) Route(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.POST("/upload/image", r.uploadFileHandler.Handler)
}

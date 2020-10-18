package router

import (
	"github.com/labstack/echo"
	"github.com/takutakukatokatojapan/image_analysis_api/api/server/handler"
)

func V1Router(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.POST("/upload/image", handler.UploadFileHandler)
}

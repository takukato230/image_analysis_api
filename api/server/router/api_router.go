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
		analysisHandler   handler.AnalysisHandler
	}
)

func NewV1Router(uploadFileHandler handler.UploadFileHandler, analysisHandler handler.AnalysisHandler) Router {
	return &V1Router{uploadFileHandler: uploadFileHandler, analysisHandler: analysisHandler}
}

func (r V1Router) Route(e *echo.Echo) {
	v1 := e.Group("/v1")
	v1.POST("/upload/file", r.uploadFileHandler.Handler)
	v1.POST("/upload/analysis", r.analysisHandler.Handler)
}

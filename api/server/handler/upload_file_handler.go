package handler

import (
	"github.com/labstack/echo"
	"github.com/takutakukatokatojapan/image_analysis_api/applications"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/appctx"
)

type (
	UploadFileHandler interface {
		Handler(c echo.Context) error
	}
	UploadFileHandlerImpl struct {
		useCase applications.UploadFileUseCase
	}
)

func NewUploadFileHandlerImpl(useCase applications.UploadFileUseCase) UploadFileHandler {
	return &UploadFileHandlerImpl{useCase: useCase}
}

func (h UploadFileHandlerImpl) Handler(c echo.Context) error {
	ctx := c.(*appctx.APPCtx)
	statusCode, res := h.useCase.Do(ctx)
	return c.JSON(statusCode, res)
}

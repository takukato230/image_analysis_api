package handler

import (
	"github.com/labstack/echo"
	"github.com/takutakukatokatojapan/image_analysis_api/applications"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/appctx"
)

type (
	AnalysisHandler interface {
		Handler(c echo.Context) error
	}
	AnalysisHandlerImpl struct {
		useCase applications.AnalysisUseCase
	}
)

func NewAnalysisHandlerImpl(useCase applications.AnalysisUseCase) AnalysisHandler {
	return &AnalysisHandlerImpl{useCase: useCase}
}

func (a AnalysisHandlerImpl) Handler(c echo.Context) error {
	ctx := c.(*appctx.APPCtx)
	statusCode, res := a.useCase.DoAnalysis(ctx)
	return c.JSON(statusCode, res)
}

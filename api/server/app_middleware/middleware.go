package app_middleware

import (
	"github.com/labstack/echo"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/appctx"
	"github.com/takutakukatokatojapan/image_analysis_api/pkg"
)

func SetUpAPPCtx(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		xRequestID := pkg.GenerateXRequestID()
		appCtx := appctx.NewAPPCtx(c, xRequestID)
		return next(appCtx)
	}
}

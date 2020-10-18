package appctx

import "github.com/labstack/echo"

type APPCtx struct {
	echo.Context
	XRequestID string
}

func NewAPPCtx(e echo.Context, xRequestID string) *APPCtx {
	return &APPCtx{
		Context:    e,
		XRequestID: xRequestID,
	}
}

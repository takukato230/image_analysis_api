package service

import (
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/appctx"
	"strings"
)

type (
	AuthService interface {
		CheckMultiPartHeader(ctx *appctx.APPCtx) bool
	}
	AuthServiceImpl struct {
	}
)

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}

func (b AuthServiceImpl) CheckMultiPartHeader(ctx *appctx.APPCtx) bool {
	acceptType := ctx.Request().Header.Get("Content-Type")
	return strings.Contains(acceptType, "multipart/form-data")
}

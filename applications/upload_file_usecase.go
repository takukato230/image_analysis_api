package applications

import (
	"fmt"
	"github.com/takutakukatokatojapan/image_analysis_api/api/types"
	"github.com/takutakukatokatojapan/image_analysis_api/domain/service"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/appctx"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/logger"
	"net/http"
)

type (
	UploadFileUseCase interface {
		Do(ctx *appctx.APPCtx) (int, types.CommonType)
	}
	UploadFileUseCaseImpl struct {
		authService        service.AuthService
		fileLoadingService service.FileLoadingService
	}
)

func NewUploadFileUseCase(authService service.AuthService, fileLoadingService service.FileLoadingService) UploadFileUseCase {
	return &UploadFileUseCaseImpl{
		authService:        authService,
		fileLoadingService: fileLoadingService,
	}
}

func (u UploadFileUseCaseImpl) Do(ctx *appctx.APPCtx) (int, types.CommonType) {
	if ok := u.authService.CheckMultiPartHeader(ctx); !ok {
		logger.Default.Warn(ctx.XRequestID, "Headerが正しくありません。")
		return http.StatusBadRequest, types.ErrorResponse{
			Code:    "E01",
			Message: "Headerが正しくありません。",
		}
	}
	container := u.fileLoadingService.Load(ctx)
	if container.Error != nil {
		logger.Default.Error(ctx.XRequestID, fmt.Sprintf("ファイルの読み込みでエラー発生: %+v", container.Error))
		return http.StatusBadRequest, types.ErrorResponse{
			Code:    "E02",
			Message: "予期せぬエラーが発生しました。",
		}
	}
	logger.Default.Debug(ctx.XRequestID, fmt.Sprintf("ファイル名: %s, コンテンツ: %s", container.FileName, container.Data))
	return http.StatusOK, types.FileUploadResponse{
		FileName:     container.FileName,
		FileContents: container.Data,
	}
}

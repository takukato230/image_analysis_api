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
	AnalysisUseCase interface {
		DoAnalysis(ctx *appctx.APPCtx) (int, types.CommonType)
	}
	AnalysisUseCaseImpl struct {
		authService        service.AuthService
		fileLoadingService service.FileLoadingService
	}
)

func NewAnalysisUseCaseImpl(authService service.AuthService, loadingService service.FileLoadingService) AnalysisUseCase {
	return &AnalysisUseCaseImpl{
		authService:        authService,
		fileLoadingService: loadingService,
	}
}

func (a AnalysisUseCaseImpl) DoAnalysis(ctx *appctx.APPCtx) (int, types.CommonType) {
	if ok := a.authService.CheckMultiPartHeader(ctx); !ok {
		logger.Default.Warn(ctx.XRequestID, "Headerが正しくありません。")
		return http.StatusBadRequest, types.ErrorResponse{
			Code:    "E01",
			Message: "Headerが正しくありません。",
		}
	}
	container := a.fileLoadingService.LoadImage(ctx)
	if container.Error != nil {
		logger.Default.Error(ctx.XRequestID, fmt.Sprintf("ファイルの読み込みでエラー発生: %+v", container.Error))
		return http.StatusInternalServerError, types.ErrorResponse{
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

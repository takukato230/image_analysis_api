package service

import (
	"fmt"
	"io/ioutil"

	"github.com/takutakukatokatojapan/image_analysis_api/domain/model"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/appctx"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/logger"
)

type (
	FileLoadingService interface {
		Load(ctx *appctx.APPCtx) FileLoadingResultContainer
		LoadImage(ctx *appctx.APPCtx) FileLoadingResultContainer
	}
	FileLoadingServiceImpl struct {
		ocr model.OCR
	}
	FileLoadingResultContainer struct {
		FileName string
		Data     string
		Error    error
	}
)

func NewFileLoadingServiceImpl(ocr model.OCR) FileLoadingService {
	return &FileLoadingServiceImpl{ocr: ocr}
}

func (f FileLoadingServiceImpl) Load(ctx *appctx.APPCtx) FileLoadingResultContainer {
	file, err := ctx.FormFile("file")
	if err != nil {
		return FileLoadingResultContainer{
			FileName: "",
			Data:     "",
			Error:    fmt.Errorf("happen where read parameter: %w", err),
		}
	}
	fileBinary, err := file.Open()
	if err != nil {
		return FileLoadingResultContainer{
			FileName: "",
			Data:     "",
			Error:    fmt.Errorf("happen where open file: %w", err),
		}
	}
	defer func() {
		if err = fileBinary.Close(); err != nil {
			logger.Default.Warn(ctx.XRequestID, fmt.Sprintf("can not close file: %v", err))
		}
	}()
	b, err := ioutil.ReadAll(fileBinary)
	if err != nil {
		return FileLoadingResultContainer{
			FileName: "",
			Data:     "",
			Error:    fmt.Errorf("happen where read file: %w", err),
		}
	}
	return FileLoadingResultContainer{
		FileName: file.Filename,
		Data:     string(b),
		Error:    nil,
	}
}

func (f FileLoadingServiceImpl) LoadImage(ctx *appctx.APPCtx) FileLoadingResultContainer {
	file, err := ctx.FormFile("image")
	if err != nil {
		return FileLoadingResultContainer{
			FileName: "",
			Data:     "",
			Error:    fmt.Errorf("happen where read parameter: %w", err),
		}
	}
	fileBinary, err := file.Open()
	if err != nil {
		return FileLoadingResultContainer{
			FileName: "",
			Data:     "",
			Error:    fmt.Errorf("happen where open file: %w", err),
		}
	}
	defer func() {
		_ = fileBinary.Close()
	}()
	b, err := ioutil.ReadAll(fileBinary)
	if err != nil {
		return FileLoadingResultContainer{
			FileName: "",
			Data:     "",
			Error:    fmt.Errorf("happen when read file: %w", err),
		}
	}
	o, err := f.ocr.Read(b)
	if err != nil {
		return FileLoadingResultContainer{
			FileName: "",
			Data:     "",
			Error:    err,
		}
	}
	return FileLoadingResultContainer{
		FileName: file.Filename,
		Data:     o,
		Error:    nil,
	}
}

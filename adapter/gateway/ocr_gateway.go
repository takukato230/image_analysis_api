package gateway

import (
	"fmt"
	"github.com/otiai10/gosseract"
	"github.com/takutakukatokatojapan/image_analysis_api/domain/model"
)

type OCRGateway struct {
}

func NewOCRGateway() model.OCR {
	return &OCRGateway{}
}

func (O OCRGateway) Read(b []byte) (string, error) {
	c := gosseract.NewClient()
	defer func() {
		_ = c.Close()
	}()
	if err := c.SetImageFromBytes(b); err != nil {
		return "", fmt.Errorf("happen when set image from bytes: %w", err)
	}
	return c.Text()
}

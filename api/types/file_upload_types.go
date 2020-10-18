package types

import "fmt"

type (
	FileUploadResponse struct {
		FileName     string `json:"file_name"`
		FileContents string `json:"file_contents"`
	}
)

func (f FileUploadResponse) OutputBody() string {
	return fmt.Sprintf("%+v", f)
}

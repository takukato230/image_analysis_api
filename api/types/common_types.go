package types

import "fmt"

type (
	CommonType interface {
		OutputBody() string
	}
	ErrorResponse struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func (e ErrorResponse) OutputBody() string {
	return fmt.Sprintf("%+v", e)
}

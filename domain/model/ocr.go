package model

type (
	OCR interface {
		Read(b []byte) (string, error)
	}
)

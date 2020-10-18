package pkg

import (
	"github.com/google/uuid"
	"time"
)

func GenerateXRequestID() string {
	var xRequestID string
	uuID, err := uuid.NewUUID()
	if err != nil {
		xRequestID = time.Now().Format("20060102150405")
	} else {
		xRequestID = uuID.String()
	}
	return xRequestID
}

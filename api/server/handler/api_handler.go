package handler

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func UploadFileHandler(c echo.Context) error {
	log.Println("hello world!!")
	return c.String(http.StatusOK, "ok")
}

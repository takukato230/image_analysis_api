package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/takutakukatokatojapan/image_analysis_api/api/server/router"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/config"
)

type (
	Server interface {
		RunServer() error
	}
	APIServer struct {
		*echo.Echo
		config config.Config
	}
)

func NewAPIServer(config config.Config) Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return &APIServer{
		Echo:   e,
		config: config,
	}
}

func (s APIServer) RunServer() error {
	router.V1Router(s.Echo)
	if err := s.Start(s.config.BindPort); err != nil {
		return err
	}
	return nil
}

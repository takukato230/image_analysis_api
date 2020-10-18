package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/takutakukatokatojapan/image_analysis_api/api/server/app_middleware"
	"github.com/takutakukatokatojapan/image_analysis_api/api/server/router"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/config"
)

type (
	Server interface {
		RunServer() error
	}
	APIServer struct {
		*echo.Echo
		config   config.Config
		v1Router router.Router
	}
)

func NewAPIServer(config config.Config, v1Router router.Router) Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(app_middleware.SetUpAPPCtx)
	return &APIServer{
		Echo:     e,
		config:   config,
		v1Router: v1Router,
	}
}

func (s APIServer) RunServer() error {
	s.v1Router.Route(s.Echo)
	if err := s.Start(s.config.BindPort); err != nil {
		return err
	}
	return nil
}

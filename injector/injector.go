package injector

import (
	"github.com/takutakukatokatojapan/image_analysis_api/adapter/gateway"
	"github.com/takutakukatokatojapan/image_analysis_api/api/server"
	"github.com/takutakukatokatojapan/image_analysis_api/api/server/handler"
	"github.com/takutakukatokatojapan/image_analysis_api/api/server/router"
	"github.com/takutakukatokatojapan/image_analysis_api/applications"
	"github.com/takutakukatokatojapan/image_analysis_api/domain/service"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/config"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/logger"
	"go.uber.org/dig"
	"log"
)

var c *dig.Container

func init() {
	c = dig.New()
	_ = c.Provide(config.NewConfig)
	// setting logger...
	if err := c.Invoke(func(config config.Config) {
		logger.NewLogger(config.APPName)
	}); err != nil {
		log.Fatal(err)
	}
	// gateway ...
	_ = c.Provide(gateway.NewOCRGateway)
	// service...
	_ = c.Provide(service.NewAuthService)
	_ = c.Provide(service.NewFileLoadingServiceImpl)
	// use case...
	_ = c.Provide(applications.NewUploadFileUseCase)
	_ = c.Provide(applications.NewAnalysisUseCaseImpl)
	// handler...
	_ = c.Provide(handler.NewUploadFileHandlerImpl)
	_ = c.Provide(handler.NewAnalysisHandlerImpl)
	// router...
	_ = c.Provide(router.NewV1Router)
	// server...
	_ = c.Provide(server.NewAPIServer)
}

func Run() error {
	if err := c.Invoke(func(s server.Server) error {
		return s.RunServer()
	}); err != nil {
		return err
	}
	return nil
}

package injector

import (
	"github.com/takutakukatokatojapan/image_analysis_api/api/server"
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/config"
	"go.uber.org/dig"
)

var c *dig.Container

func init() {
	c = dig.New()
	_ = c.Provide(config.NewConfig)
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

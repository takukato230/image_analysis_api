package injector

import (
	"github.com/takutakukatokatojapan/image_analysis_api/infrastructure/config"
	"go.uber.org/dig"
)

var c = dig.New()

func init() {
	c.Provide(config.NewConfig)
}

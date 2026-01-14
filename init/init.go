package init

import (
	"github.com/go-xuan/elasticx"

	"github.com/go-xuan/quanx/configx"
	"github.com/go-xuan/utilx/errorx"
)

func init() {
	errorx.Panic(Init())
}

func Init() error {
	var err error
	if err = configx.LoadConfigurator(&elasticx.Configs{}); err == nil && elasticx.Initialized() {
		return nil
	} else if err = configx.LoadConfigurator(&elasticx.Config{}); err == nil && elasticx.Initialized() {
		return nil
	}
	return errorx.Wrap(err, "init elastic search failed")
}

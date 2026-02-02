package elasticx

import (
	"github.com/go-xuan/configx"
	log "github.com/sirupsen/logrus"
)

func init() {
	Init() // 初始化 ES
}

func Init() {
	logger := log.WithField("package", "elasticx")
	if err := configx.LoadConfigurator(&Configs{}); err == nil && Initialized() {
		logger.Info("initialized success")
		return
	}
	if err := configx.LoadConfigurator(&Config{}); err == nil && Initialized() {
		logger.Info("initialized success")
		return
	}
	logger.Warn("initialized failed")
}

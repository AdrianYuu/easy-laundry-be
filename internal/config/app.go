package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	Viper *viper.Viper
	Fiber *fiber.App
	Log   *logrus.Logger
	Db    *gorm.DB
}

func Bootstrap(config *BootstrapConfig) {

}

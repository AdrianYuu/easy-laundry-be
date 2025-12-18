package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	Config *viper.Viper
	App    *fiber.App
}

func Bootstrap(config *BootstrapConfig) {

}

package config

import (
	"github.com/AdrianYuu/easy-laundry-be/internal/delivery/http"
	"github.com/AdrianYuu/easy-laundry-be/internal/delivery/http/route"
	"github.com/AdrianYuu/easy-laundry-be/internal/repository"
	"github.com/AdrianYuu/easy-laundry-be/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	Viper    *viper.Viper
	Fiber    *fiber.App
	Log      *logrus.Logger
	Db       *gorm.DB
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {
	userRepository := repository.NewUserRepository(config.Log)
	userUseCase := usecase.NewUserUseCase(config.Db, config.Log, config.Validate, userRepository)
	userController := http.NewUserController(config.Log, userUseCase)

	routeConfig := route.RouteConfig{
		Fiber:          config.Fiber,
		UserController: userController,
	}
	routeConfig.Setup()
}

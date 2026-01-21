package http

import (
	"github.com/AdrianYuu/easy-laundry-be/internal/model/response"
	"github.com/AdrianYuu/easy-laundry-be/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log         *logrus.Logger
	UserUseCase *usecase.UserUseCase
}

func NewUserController(log *logrus.Logger, userUseCase *usecase.UserUseCase) *UserController {
	return &UserController{
		Log:         log,
		UserUseCase: userUseCase,
	}
}

func (c *UserController) Ping(ctx *fiber.Ctx) error {
	return ctx.JSON(response.WebResponse[string]{
		Data: "Pong",
	})
}

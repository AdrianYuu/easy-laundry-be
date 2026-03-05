package usecase

import (
	"github.com/AdrianYuu/easy-laundry-be/internal/helper"
	"github.com/AdrianYuu/easy-laundry-be/internal/model/request"
	"github.com/AdrianYuu/easy-laundry-be/internal/model/response"
	"github.com/AdrianYuu/easy-laundry-be/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserUseCase struct {
	Db             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	Viper          *viper.Viper
	UserRepository *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, viper *viper.Viper, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		Db:             db,
		Log:            log,
		Validate:       validate,
		Viper:          viper,
		UserRepository: userRepository,
	}
}

func (c *UserUseCase) Verify(request *request.VerifyRequest) (*response.VerifyResponse, error) {
	err := c.Validate.Struct(request)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	claims, err := helper.ExtractJwt(request.Token, c.Viper)
	if err != nil {
		return nil, err
	}

	return &response.VerifyResponse{
		Id: claims.UserId,
	}, nil
}

package middleware

import (
	"github.com/AdrianYuu/easy-laundry-be/internal/model/request"
	"github.com/AdrianYuu/easy-laundry-be/internal/model/response"
	"github.com/AdrianYuu/easy-laundry-be/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware(userUseCase *usecase.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenString := ctx.Get("Authorization", "NOT_FOUND")
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		auth, err := userUseCase.Verify(
			&request.VerifyRequest{
				Token: tokenString,
			},
		)
		if err != nil {
			return err
		}

		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetAuth(ctx *fiber.Ctx) *response.VerifyResponse {
	return ctx.Locals("auth").(*response.VerifyResponse)
}

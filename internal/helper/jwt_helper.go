package helper

import (
	"github.com/AdrianYuu/easy-laundry-be/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJwt(userId string, viper *viper.Viper) (string, error) {
	claims := model.JwtClaim{
		UserId:           userId,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "failed to generate access token")
	}

	return signed, nil
}

func ExtractJwt(tokenString string, viper *viper.Viper) (*model.JwtClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid signing method")
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})

	if err != nil || !token.Valid {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid or expired token")
	}

	claims, ok := token.Claims.(*model.JwtClaim)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "invalid token claims")
	}

	return claims, nil
}

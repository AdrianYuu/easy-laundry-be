package model

import "github.com/golang-jwt/jwt/v5"

type JwtClaim struct {
	UserId string
	jwt.RegisteredClaims
}

package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId   string
	TenantId string
	jwt.StandardClaims
}

var JWTKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWTToken(userId string) (string, error) {
	expTime := time.Now().Add(time.Hour * 2).Unix()

	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "tenant-onboarding",
			Subject:   userId,
			ExpiresAt: expTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTKey)
}

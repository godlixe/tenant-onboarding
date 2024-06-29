package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"tenant-onboarding/pkg/auth"
	"tenant-onboarding/pkg/framework"
	"tenant-onboarding/pkg/httpx"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type IAMClaims struct {
	Email      string `json:"email,omitempty"`
	Name       string `json:"name,omitempty"`
	PictureUrl string `json:"picture,omitempty"`
	jwt.StandardClaims
}

func Authenticate() gin.HandlerFunc {

	fmt.Println(framework.CheckIntegratedMode(), "integration mode")
	// On integrated mode, use external IAM API
	// for authentication and authorization
	if framework.CheckIntegratedMode() {

		return func(ctx *gin.Context) {
			token := ctx.GetHeader("Authorization")
			if token == "" {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
					Message: "unauthorized user",
				})
				return
			}

			tokenHeader := strings.Split(token, "Bearer ")

			if len(tokenHeader) < 2 {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
					Message: "corrupt header format",
				})
				return
			}

			claims := IAMClaims{}
			jwtToken, err := jwt.ParseWithClaims(tokenHeader[1], &claims, func(token *jwt.Token) (interface{}, error) {
				return auth.JWTKey, nil
			})

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
					Message: "error parsing token",
				})
				return
			}

			if !jwtToken.Valid {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
					Message: "invalid token",
				})
				return
			}

			ctx.Set("token", jwtToken.Raw)
			ctx.Set("user_id", claims.Subject)
			ctx.Next()
		}
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
				Message: "unauthorized user",
			})
			return
		}

		tokenHeader := strings.Split(token, "Bearer ")

		if len(tokenHeader) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
				Message: "corrupt header format",
			})
			return
		}

		claims := auth.Claims{}
		jwtToken, err := jwt.ParseWithClaims(tokenHeader[1], &claims, func(token *jwt.Token) (interface{}, error) {
			return auth.JWTKey, nil
		})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
				Message: "error parsing token",
			})
			return
		}

		if !jwtToken.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, httpx.Response{
				Message: "invalid token",
			})
			return
		}

		ctx.Set("user_id", claims.UserId)
		ctx.Next()
	}
}

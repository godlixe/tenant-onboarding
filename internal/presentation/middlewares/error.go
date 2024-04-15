package middlewares

import (
	"net/http"
	"tenant-onboarding/pkg/httpx"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ErrorHandler(logger zerolog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			switch e := err.Err.(type) {
			case *httpx.HttpError:
				ctx.JSON(e.StatusCode, httpx.Response{
					Message: e.Message,
				})
				log.Error().Str("message", e.Error())
			default:
				ctx.JSON(http.StatusInternalServerError, httpx.Response{
					Message: e.Error(),
				})
				log.Error().Str("message", e.Error())
			}
		}

		ctx.Abort()
	}
}

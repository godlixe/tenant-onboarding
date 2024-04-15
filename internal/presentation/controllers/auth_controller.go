package controllers

import (
	"net/http"
	"tenant-onboarding/internal/domain/users/entity"
	"tenant-onboarding/internal/presentation/dto"
	"tenant-onboarding/internal/presentation/services"
	"tenant-onboarding/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(
	authService services.AuthService,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var params dto.LoginDTO

	err := ctx.ShouldBind(&params)
	if err != nil {
		err = httpx.NewError("validation error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	res, err := c.authService.Login(ctx, &entity.User{
		Username: params.Username,
		Password: params.Password,
	})
	if err != nil {
		err = httpx.NewError("error logging in", err, http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "login successfull",
		Data: struct {
			Token string `json:"token"`
		}{
			Token: res,
		},
	})
}

func (c *AuthController) Register(ctx *gin.Context) {
	var params dto.RegisterDTO

	err := ctx.ShouldBind(&params)
	if err != nil {
		err = httpx.NewError("validation error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	res, err := c.authService.Register(
		ctx,
		&entity.User{
			Name:     params.Name,
			Username: params.Username,
			Email:    params.Email,
			Password: params.Password,
		},
	)
	if err != nil {
		err = httpx.NewError("error registering", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "register successfull",
		Data:    res,
	})
}

func (c *AuthController) Me(ctx *gin.Context) {
	res, err := c.authService.Me(
		ctx,
	)
	if err != nil {
		err = httpx.NewError("error getting me", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "register successfull",
		Data:    res,
	})
}

package controllers

import (
	"net/http"
	"tenant-onboarding/modules/auth/internal/app/commands"
	"tenant-onboarding/modules/auth/internal/app/queries"
	"tenant-onboarding/modules/auth/internal/presentation/dto"
	"tenant-onboarding/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserRegister *commands.UserRegisterCommand
	UserLogin    *commands.UserLoginCommand

	UserQuery queries.UserQuery
}

func NewAuthController(
	UserRegister *commands.UserRegisterCommand,
	UserLogin *commands.UserLoginCommand,
	UserQuery queries.UserQuery,
) *AuthController {
	return &AuthController{
		UserRegister: UserRegister,
		UserLogin:    UserLogin,
		UserQuery:    UserQuery,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var params dto.RegisterDTO

	err := ctx.ShouldBind(&params)
	if err != nil {
		err = httpx.NewError("validation error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	req := commands.NewUserRegisterRequest(
		params.Name,
		params.Email,
		params.Username,
		params.Password,
	)

	err = c.UserRegister.Execute(ctx, req)
	if err != nil {
		err = httpx.NewError("error registering", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "register successfull",
		Data:    nil,
	})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var params dto.LoginDTO

	err := ctx.ShouldBind(&params)
	if err != nil {
		err = httpx.NewError("validation error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	req := commands.NewUserLoginRequest(
		params.Email,
		params.Password,
	)

	token, err := c.UserLogin.Execute(ctx, req)
	if err != nil {
		err = httpx.NewError("error login", err, http.StatusUnauthorized)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "login successfull",
		Data:    token,
	})
}

func (c *AuthController) Me(ctx *gin.Context) {
	userID := ctx.Value("user_id").(string)
	user, err := c.UserQuery.GetByID(ctx, userID)
	if err != nil {
		err = httpx.NewError("get user error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "login successfull",
		Data:    user,
	})
}

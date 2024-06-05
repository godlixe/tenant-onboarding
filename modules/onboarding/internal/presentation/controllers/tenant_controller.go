package controllers

import (
	"net/http"
	"tenant-onboarding/modules/onboarding/internal/app/commands"
	"tenant-onboarding/modules/onboarding/internal/presentation/dto"
	"tenant-onboarding/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type TenantController struct {
	userCreateTenant *commands.UserCreateTenantCommand
}

func NewTenantController(
	userCreateTenant *commands.UserCreateTenantCommand,
) *TenantController {
	return &TenantController{
		userCreateTenant: userCreateTenant,
	}
}

func (c *TenantController) CreateTenant(ctx *gin.Context) {
	var params dto.CreateTenantDTO

	err := ctx.ShouldBind(&params)
	if err != nil {
		err = httpx.NewError("validation error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	req := commands.NewUserCreateTenantRequest(
		params.ProductID,
		params.OrganizationID,
		params.Name,
	)
	err = c.userCreateTenant.Execute(
		ctx,
		req,
	)
	if err != nil {
		err = httpx.NewError("error creating tenant", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "tenant successfully created",
		Data:    nil,
	})
}

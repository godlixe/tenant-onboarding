package controllers

import (
	"net/http"
	"tenant-onboarding/internal/domain/users/entity"
	"tenant-onboarding/internal/presentation/dto"
	"tenant-onboarding/internal/presentation/services"
	"tenant-onboarding/pkg/httpx"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TenantController struct {
	TenantService services.TenantService
}

func NewTenantController(
	TenantService services.TenantService,
) *TenantController {
	return &TenantController{
		TenantService: TenantService,
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

	res, err := c.TenantService.CreateTenant(
		ctx,
		&entity.Tenant{
			ProductID: uuid.MustParse(params.ProductID),
			Name:      params.Name,
			Subdomain: params.Subdomain,
			Status:    entity.TenantCreated,
		},
	)
	if err != nil {
		err = httpx.NewError("error creating tenant", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "tenant successfully created",
		Data:    res,
	})
}

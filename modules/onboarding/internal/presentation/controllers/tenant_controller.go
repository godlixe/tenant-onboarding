package controllers

import (
	"net/http"
	"tenant-onboarding/modules/onboarding/internal/app/commands"
	"tenant-onboarding/modules/onboarding/internal/app/queries"
	"tenant-onboarding/modules/onboarding/internal/presentation/dto"
	"tenant-onboarding/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type TenantController struct {
	userCreateTenant *commands.UserCreateTenantCommand
	tenantQuery      queries.TenantQuery
}

func NewTenantController(
	userCreateTenant *commands.UserCreateTenantCommand,
	tenantQuery queries.TenantQuery,
) *TenantController {
	return &TenantController{
		userCreateTenant: userCreateTenant,
		tenantQuery:      tenantQuery,
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
	tenant, err := c.userCreateTenant.Execute(
		ctx,
		req,
	)
	if err != nil {
		err = httpx.NewError("error creating tenant", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	type TenantResponse struct {
		ID             string `json:"id"`
		OrganizationID string `json:"organization_id"`
		ProductID      string `json:"product_id"`
		Status         string `json:"status"`
		Name           string `json:"name"`
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "tenant successfuly created",
		Data: TenantResponse{
			ID:             tenant.ID.String(),
			OrganizationID: tenant.OrganizationID.String(),
			ProductID:      tenant.ProductID.String(),
			Status:         tenant.Status.String(),
			Name:           tenant.Name,
		},
	})
}

func (c *TenantController) GetTenants(ctx *gin.Context) {
	organizationID := ctx.Query("organization_id")

	tenants, err := c.tenantQuery.GetAllByOrganizationID(ctx, organizationID)
	if err != nil {
		err = httpx.NewError("get tenants error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "get tenants successful",
		Data:    tenants,
	})
}

package controllers

import (
	"net/http"
	"tenant-onboarding/modules/auth/internal/app/commands"
	"tenant-onboarding/modules/auth/internal/app/queries"
	"tenant-onboarding/modules/auth/internal/presentation/dto"
	"tenant-onboarding/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type OrganizationController struct {
	UserCreateOrganization *commands.UserCreateOrganizationCommand
	OrganizationQuery      queries.OrganizationQuery
}

func NewOrganizationController(
	UserCreateOrganization *commands.UserCreateOrganizationCommand,
	OrganizationQuery queries.OrganizationQuery,
) *OrganizationController {
	return &OrganizationController{
		UserCreateOrganization: UserCreateOrganization,
		OrganizationQuery:      OrganizationQuery,
	}
}

func (c *OrganizationController) GetAllUserOrganization(ctx *gin.Context) {
	userID := ctx.Value("user_id").(string)

	orgs, err := c.OrganizationQuery.GetAllUserOrganization(ctx, userID)
	if err != nil {
		err = httpx.NewError("error getting organizations", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "get organizations successfull",
		Data:    orgs,
	})
}

func (c *OrganizationController) CreateOrganizations(ctx *gin.Context) {
	userID := ctx.Value("user_id").(string)

	var params dto.CreateOrganizationDTO

	err := ctx.ShouldBind(&params)
	if err != nil {
		err = httpx.NewError("validation error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	req := commands.NewUserCreateOrganizationRequest(
		params.Name,
		params.Subdomain,
		userID,
	)

	err = c.UserCreateOrganization.Execute(ctx, req)
	if err != nil {
		err = httpx.NewError("error creating organization", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "create organization successfull",
		Data:    nil,
	})
}

package controllers

import (
	"net/http"
	"tenant-onboarding/modules/onboarding/internal/app/queries"
	"tenant-onboarding/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	appQuery queries.AppQuery
}

func NewAppController(
	appQuery queries.AppQuery,
) *AppController {
	return &AppController{
		appQuery: appQuery,
	}
}

func (c *AppController) GetAll(ctx *gin.Context) {
	// var filters queries.AppFilter

	// err := ctx.BindQuery(&filters)
	// if err != nil {
	// 	err = httpx.NewError("binding error", err, http.StatusBadRequest)
	// 	ctx.Error(err)
	// 	return

	apps, err := c.appQuery.GetAll(ctx)
	if err != nil {
		err = httpx.NewError("get Apps error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "get Apps successful",
		Data:    apps,
	})
}

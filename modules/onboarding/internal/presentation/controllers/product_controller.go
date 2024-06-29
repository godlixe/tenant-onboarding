package controllers

import (
	"net/http"
	"tenant-onboarding/modules/onboarding/internal/app/queries"
	"tenant-onboarding/pkg/httpx"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductQuery queries.ProductQuery
}

func NewProductController(
	ProductQuery queries.ProductQuery,
) *ProductController {
	return &ProductController{
		ProductQuery: ProductQuery,
	}
}

func (c *ProductController) GetAll(ctx *gin.Context) {
	var filters queries.ProductFilter

	err := ctx.BindQuery(&filters)
	if err != nil {
		err = httpx.NewError("binding error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	products, err := c.ProductQuery.GetProducts(ctx, &filters)
	if err != nil {
		err = httpx.NewError("get products error", err, http.StatusBadRequest)
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, httpx.Response{
		Message: "get products successful",
		Data:    products,
	})
}

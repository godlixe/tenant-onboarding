package dto

type CreateTenantDTO struct {
	ProductID string `json:"product_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Subdomain string `json:"subdomain" binding:"required"`
}

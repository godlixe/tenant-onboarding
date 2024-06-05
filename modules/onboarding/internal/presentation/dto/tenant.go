package dto

type CreateTenantDTO struct {
	ProductID      string `json:"product_id" binding:"required"`
	OrganizationID string `json:"organization_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
}

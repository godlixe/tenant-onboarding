package dto

type CreateOrganizationDTO struct {
	Name       string `json:"name" binding:"required"`
	Identifier string `json:"identifier" binding:"required"`
}

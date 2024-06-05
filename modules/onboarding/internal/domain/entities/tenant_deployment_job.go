package entities

import "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"

type TenantData struct {
	ID             string `json:"id"`
	ProductID      string `json:"product_id"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
	Status         string `json:"status"`
}

type InfrastructureData struct {
	ID              string                       `json:"id"`
	ProductID       string                       `json:"product_id"`
	Name            string                       `json:"name"`
	DeploymentModel valueobjects.DeploymentModel `json:"deployment_model"`
	UserCount       int                          `json:"user_count"`
	UserLimit       int                          `json:"user_limit"`
	Metadata        string                       `json:"metadata"`
}

type TenantDeploymentJob struct {
	TenantData  TenantData  `json:"tenant_data"`
	ProductData ProductData `json:"product_data"`
}

type ProductData struct {
	ID               string `json:"id"`
	AppID            int    `json:"app_id"`
	TierID           int    `json:"tier_id"`
	DeploymentSchema string `json:"deployment_schema"`
}

func NewTenantDeploymentJob(
	tenant *Tenant,
	Product *Product,
) *TenantDeploymentJob {

	tenantData := TenantData{
		ID:             tenant.ID.String(),
		ProductID:      tenant.ProductID.String(),
		OrganizationID: tenant.OrganizationID.String(),
		Name:           tenant.Name,
		Status:         tenant.Status.String(),
	}

	productData := ProductData{
		ID:               Product.ID.String(),
		AppID:            Product.AppID,
		TierID:           Product.TierID,
		DeploymentSchema: Product.DeploymentSchema,
	}

	return &TenantDeploymentJob{
		TenantData:  tenantData,
		ProductData: productData,
	}
}

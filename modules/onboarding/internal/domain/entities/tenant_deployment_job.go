package entities

type TenantDeploymentJob struct {
	TenantID  string `json:"tenant_id"`
	ProductID string `json:"product_id"`
}

func NewTenantDeploymentJob(
	tenant *Tenant,
	product *Product,
) *TenantDeploymentJob {
	return &TenantDeploymentJob{
		TenantID:  tenant.ID.String(),
		ProductID: product.ID.String(),
	}
}

package entities

type TenantDeploymentJob struct {
	TenantID string `json:"tenant_id"`
}

func NewTenantDeploymentJob(
	tenant *Tenant,
) *TenantDeploymentJob {
	return &TenantDeploymentJob{
		TenantID: tenant.ID.String(),
	}
}

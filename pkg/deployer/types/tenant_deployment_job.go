package types

// TenantDeploymentJob defines a minimal tenant deployment job data.
type TenantDeploymentJob struct {
	TenantID  string `json:"tenant_id"`
	ProductID string `json:"product_id"`
}

package errors

import "errors"

var (
	ErrInvalidAppID            = errors.New("invalid_app_id")
	ErrInvalidTierID           = errors.New("invalid_tier_id")
	ErrInvalidInfrastructureID = errors.New("invalid_infrastructure_id")
	ErrInvalidUserID           = errors.New("invalid_user_id")
	ErrInvalidTenantID         = errors.New("invalid_tenant_id")
	ErrInvalidProjectID        = errors.New("invalid_project_id")
	ErrInvalidProductID        = errors.New("invalid_product_id")
	ErrInvalidOrganizationID   = errors.New("invalid_organization_id")

	ErrInvalidDeploymentModel    = errors.New("invalid_deployment_model")
	ErrInvalidTenantStatus       = errors.New("invalid_tenant_status")
	ErrInvalidInfrastructureType = errors.New("invalid_infrastructure_type")
)

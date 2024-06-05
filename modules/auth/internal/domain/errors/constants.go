package errors

import "errors"

var (
	ErrInvalidUserID         = errors.New("invalid_user_id")
	ErrInvalidTenantID       = errors.New("invalid_tenant_id")
	ErrInvalidProjectID      = errors.New("invalid_project_id")
	ErrInvalidProductID      = errors.New("invalid_product_id")
	ErrInvalidOrganizationID = errors.New("invalid_organization_id")

	ErrInvalidDeploymentModel = errors.New("invalid_deployment_model")
	ErrInvalidTenantStatus    = errors.New("invalid_tenant_status")

	ErrEmailExists        = errors.New("duplicate_user_email")
	ErrOrgSubdomainExists = errors.New("duplicate_organization_subdomain")

	ErrUnauthenticated         = errors.New("user_unauthenticated")
	ErrInvalidOrganizationRole = errors.New("invalid_organization_role")
)

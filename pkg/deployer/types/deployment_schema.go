package types

// DeploymentSchema defines the deployment schema of
// a product. A deployment schema of a product is predefined.
type DeploymentSchema struct {
	// Github repository of the directory for deployment.
	DeploymentRepositoryURL string `json:"deployment_repository_url"`
	// Terraform execution path.
	TerraformExecutionPath string `json:"terraform_execution_path"`
	// Deployment model.
	DeploymentModel DeploymentModel `json:"deployment_model"`
	// Custom init script execution path.
	InitScriptPath string `json:"init_script_path"`
	// Custom migration script execution path.
	MigrationScriptPath string `json:"migration_script_path"`

	// Product identifier
	ProductID string
}

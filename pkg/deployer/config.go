package deployer

// Config holds the configuration of the Terraform Deployer.
// Config is customizable to tailor specific needs.
type Config struct {
	// Absolute path of google service account.
	GoogleServiceAccountAbsolutePath string

	// Project ID of Google Project.
	GoogleProjectID string

	// Region of Google deployment.
	GoogleDeploymentRegion string

	// Absolute path of terraform folder in project.
	TerraformAbsolutePath string

	// Absolute path of terraform executable.
	TerraformExecPath string

	// Terraform backend bucket name.
	TerraformBackendBucket string
}

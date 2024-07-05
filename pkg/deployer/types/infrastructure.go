package types

type InfraOutput struct {
	Metadata             string
	ResourceInformations string
}

// RawInfrastructure is a placeholder struct
// for the infrastructure that will be deployed by the pipeline.
type RawInfrastructure struct {
	ID          string
	Metadata    *InfraOutput
	IsCreateNew bool
}

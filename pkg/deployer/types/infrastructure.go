package types

// RawInfrastructure is a placeholder struct
// for the infrastructure that will be deployed by the pipeline.
type RawInfrastructure struct {
	ID          string
	Metadata    string
	IsCreateNew bool
}

package types

import "encoding/json"

type InfraOutput struct {
	Metadata             json.RawMessage `json:"metadata"`
	ResourceInformations json.RawMessage `json:"resource_information"`
}

// RawInfrastructure is a placeholder struct
// for the infrastructure that will be deployed by the pipeline.
type RawInfrastructure struct {
	ID          string
	Metadata    *InfraOutput
	IsCreateNew bool
}

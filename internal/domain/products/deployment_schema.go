package products

import "errors"

// DeploymentSchema contains information
// of the deployment schema of a resource. The field
// is represented in JSON which allows dynamic
// changes to the structure.
type DeploymentSchema struct {
	Web     WebSchema     `json:"web"`
	Compute ComputeSchema `json:"compute"`
	Storage StorageSchema `json:"storage"`
}

type WebSchema struct {
	Model  DeploymentModel
	Source string
}

type ComputeSchema struct {
	Model  DeploymentModel
	Source string
}

type StorageSchema struct {
	Model  DeploymentModel
	Source string
}

// DeploymentModel defines a deployment model
// of a resource. There are 2 deployment models
// supported, silo and pool.
type DeploymentModel struct {
	model string
}

var (
	Silo = DeploymentModel{"silo"}
	Pool = DeploymentModel{"pool"}
)

func (d DeploymentModel) String() string {
	return d.model
}

func FromString(s string) (DeploymentModel, error) {
	switch s {
	case Silo.model:
		return Silo, nil
	case Pool.model:
		return Pool, nil
	default:
		return DeploymentModel{}, errors.New("invalid model")
	}
}

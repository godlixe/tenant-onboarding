package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

var ErrInvalidDeploymentModel = errors.New("invalid_deployment_model")

// DeploymentModel defines a deployment model
// of a resource. There are 2 deployment models
// supported, silo and pool.
type DeploymentModel struct {
	Model string
}

var (
	Silo = DeploymentModel{"silo"}
	Pool = DeploymentModel{"pool"}
)

func (d DeploymentModel) String() string {
	return d.Model
}

func NewDeploymentModel(s string) (DeploymentModel, error) {
	switch s {
	case Silo.Model:
		return Silo, nil
	case Pool.Model:
		return Pool, nil
	default:
		return DeploymentModel{}, ErrInvalidDeploymentModel
	}
}

func (d *DeploymentModel) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		res, err := NewDeploymentModel(v)
		if err != nil {
			return err
		}

		d.Model = res.Model
	default:
		return ErrInvalidDeploymentModel
	}

	return nil
}

func (d *DeploymentModel) Value() (driver.Value, error) {
	if d == nil {
		return nil, nil
	}

	return d.Model, nil
}

func (d *DeploymentModel) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	var deploymentModelValue string
	err := json.Unmarshal(data, &deploymentModelValue)
	if err != nil {
		return err
	}

	deploymentModel, err := NewDeploymentModel(deploymentModelValue)
	if err != nil {
		return err
	}

	d.Model = deploymentModel.Model
	return nil
}

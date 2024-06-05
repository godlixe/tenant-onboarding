package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// ResourceInformation contains information
// of the resources a tenant use. The field
// is represented in JSON which allows dynamic
// changes to the structure.
type ResourceInformation struct {
	Gateway GatewayInformation `json:"gateway"`
	Web     WebInformation     `json:"web"`
	Compute ComputeInformation `json:"compute"`
	Storage StorageInformation `json:"storages"`
}

type GatewayInformation struct {
	URL string `json:"url"`
}

type WebInformation struct {
	URL string `json:"url"`
}

type ComputeInformation struct {
	URL string `json:"url"`
}

type StorageInformation struct {
	URL string `json:"url"`
}

func (r *ResourceInformation) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("failed to parse value")
	}

	err := json.Unmarshal([]byte(str), r)
	if err != nil {
		return err
	}

	return nil
}

func (r *ResourceInformation) Value() (driver.Value, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return string(b), nil
}

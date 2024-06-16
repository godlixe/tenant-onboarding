package valueobjects

import (
	"database/sql/driver"
	"tenant-onboarding/modules/onboarding/internal/errorx"

	"github.com/google/uuid"
)

var ErrInvalidInfrastructureID = errorx.ErrInvalidInfrastructureID

type InfrastructureID struct {
	ID string
}

func NewInfrastructureID(id string) (InfrastructureID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return InfrastructureID{}, ErrInvalidProductID
	}

	return InfrastructureID{id}, nil
}

func GenerateInfrastructureID() InfrastructureID {
	return InfrastructureID{uuid.NewString()}
}

func (i InfrastructureID) Equals(other InfrastructureID) bool {
	return i.ID == other.ID
}

func (i InfrastructureID) String() string {
	return i.ID
}

func (i *InfrastructureID) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		res, err := NewInfrastructureID(v)
		if err != nil {
			return err
		}

		i.ID = res.ID
	default:
		return ErrInvalidInfrastructureID
	}

	return nil
}

func (i *InfrastructureID) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return i.ID, nil
}

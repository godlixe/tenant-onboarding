package valueobjects

import (
	"database/sql/driver"
	"strings"
	"tenant-onboarding/modules/onboarding/internal/errors"

	"github.com/google/uuid"
)

var ErrInvalidOrganizationID = errors.ErrInvalidOrganizationID

type OrganizationID struct {
	ID string
}

func NewOrganizationID(id string) (OrganizationID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return OrganizationID{}, ErrInvalidOrganizationID
	}

	return OrganizationID{id}, nil
}

func GenerateOrganizationID() OrganizationID {
	return OrganizationID{uuid.NewString()}
}

func (i OrganizationID) String() string {
	return i.ID
}

func (i OrganizationID) Equals(other OrganizationID) bool {
	return strings.EqualFold(i.ID, other.ID)
}

func (i *OrganizationID) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		res, err := NewOrganizationID(v)
		if err != nil {
			return err
		}

		i.ID = res.ID
	default:
		return ErrInvalidOrganizationID
	}

	return nil
}

func (i *OrganizationID) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return i.ID, nil
}

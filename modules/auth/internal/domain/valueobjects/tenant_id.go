package valueobjects

import (
	"database/sql/driver"
	"strings"
	"tenant-onboarding/modules/auth/internal/domain/errors"

	"github.com/google/uuid"
)

var ErrInvalidTenantId = errors.ErrInvalidTenantID

type TenantID struct {
	ID string
}

func NewTenantID(id string) (TenantID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return TenantID{}, ErrInvalidTenantId
	}

	return TenantID{uuid.NewString()}, nil
}

func GenerateTenantID() TenantID {
	return TenantID{uuid.NewString()}
}

func (i TenantID) String() string {
	return i.ID
}

func (i TenantID) Equals(other TenantID) bool {
	return strings.EqualFold(i.ID, other.ID)
}

func (i *TenantID) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		res, err := NewTenantID(v)
		if err != nil {
			return err
		}

		i.ID = res.ID
	default:
		return ErrInvalidTenantId
	}

	return nil
}

func (i *TenantID) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return i.ID, nil
}

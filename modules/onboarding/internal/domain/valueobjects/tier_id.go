package valueobjects

import (
	"database/sql/driver"
	"tenant-onboarding/modules/onboarding/internal/errors"
)

var ErrInvalidTierID = errors.ErrInvalidTierID

type TierID struct {
	ID int
}

func NewTierID(id int) (TierID, error) {
	return TierID{id}, nil
}

func (i TierID) Equals(other TierID) bool {
	return i.ID == other.ID
}

func (i TierID) Int() int {
	return i.ID
}

func (i *TierID) Scan(value interface{}) error {
	switch v := value.(type) {
	case int:
		res, err := NewTierID(v)
		if err != nil {
			return err
		}

		i.ID = res.ID
	default:
		return ErrInvalidTierID
	}

	return nil
}

func (i *TierID) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return i.ID, nil
}

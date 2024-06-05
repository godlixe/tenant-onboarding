package valueobjects

import (
	"database/sql/driver"
	"tenant-onboarding/modules/onboarding/internal/errors"
)

var ErrInvalidAppID = errors.ErrInvalidAppID

type AppID struct {
	ID int
}

func NewAppID(id int) (AppID, error) {
	return AppID{id}, nil
}

func (i AppID) Equals(other AppID) bool {
	return i.ID == other.ID
}

func (i AppID) Int() int {
	return i.ID
}

func (i *AppID) Scan(value interface{}) error {
	switch v := value.(type) {
	case int:
		res, err := NewAppID(v)
		if err != nil {
			return err
		}

		i.ID = res.ID
	default:
		return ErrInvalidAppID
	}

	return nil
}

func (i *AppID) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return i.ID, nil
}

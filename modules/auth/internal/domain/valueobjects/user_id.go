package valueobjects

import (
	"database/sql/driver"
	"strings"
	"tenant-onboarding/modules/auth/internal/domain/errors"

	"github.com/google/uuid"
)

var ErrInvalidUserId = errors.ErrInvalidUserID

type UserID struct {
	ID string
}

func NewUserID(id string) (UserID, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return UserID{}, ErrInvalidUserId
	}

	return UserID{id}, nil
}

func GenerateUserID() UserID {
	return UserID{uuid.NewString()}
}

func (i UserID) String() string {
	return i.ID
}

func (i UserID) Equals(other UserID) bool {
	return strings.EqualFold(i.ID, other.ID)
}

func (i *UserID) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		res, err := NewUserID(v)
		if err != nil {
			return err
		}

		i.ID = res.ID
	default:
		return ErrInvalidUserId
	}

	return nil
}

func (i *UserID) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return i.ID, nil
}

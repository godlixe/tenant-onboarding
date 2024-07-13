package valueobjects

import (
	"database/sql/driver"
	"strings"
	"tenant-onboarding/modules/auth/internal/domain/errors"

	"github.com/google/uuid"
)

var ErrInvalidOrganizationLevel = errors.ErrInvalidOrganizationLevel

type OrganizationLevel struct {
	Level string
}

var (
	LevelMember  = OrganizationLevel{"member"}
	LevelManager = OrganizationLevel{"manager"}
	LevelOwner   = OrganizationLevel{"owner"}
)

func NewOrganizationLevel(level string) (OrganizationLevel, error) {
	switch level {
	case LevelMember.Level:
		return LevelMember, nil
	case LevelOwner.Level:
		return LevelOwner, nil
	}

	return OrganizationLevel{}, ErrInvalidOrganizationLevel
}

func GenerateOrganizationLevel() OrganizationLevel {
	return OrganizationLevel{uuid.NewString()}
}

func (r OrganizationLevel) String() string {
	return r.Level
}

func (r OrganizationLevel) Equals(other OrganizationLevel) bool {
	return strings.EqualFold(r.Level, other.Level)
}

func (r *OrganizationLevel) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		res, err := NewOrganizationLevel(v)
		if err != nil {
			return err
		}

		r.Level = res.Level
	default:
		return ErrInvalidOrganizationLevel
	}

	return nil
}

func (r *OrganizationLevel) Value() (driver.Value, error) {
	if r == nil {
		return nil, nil
	}

	return r.Level, nil
}

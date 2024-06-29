package valueobjects

import (
	"database/sql/driver"
	"strings"
	"tenant-onboarding/modules/auth/internal/domain/errors"

	"github.com/google/uuid"
)

var ErrInvalidOrganizationRole = errors.ErrInvalidOrganizationRole

type OrganizationRole struct {
	Role string
}

var (
	RoleMember  = OrganizationRole{"member"}
	RoleManager = OrganizationRole{"manager"}
	RoleOwner   = OrganizationRole{"owner"}
)

func NewOrganizationRole(role string) (OrganizationRole, error) {
	switch role {
	case RoleMember.Role:
		return RoleMember, nil
	case RoleOwner.Role:
		return RoleOwner, nil
	}

	return OrganizationRole{}, ErrInvalidOrganizationRole
}

func GenerateOrganizationRole() OrganizationRole {
	return OrganizationRole{uuid.NewString()}
}

func (r OrganizationRole) String() string {
	return r.Role
}

func (r OrganizationRole) Equals(other OrganizationRole) bool {
	return strings.EqualFold(r.Role, other.Role)
}

func (r *OrganizationRole) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		res, err := NewOrganizationRole(v)
		if err != nil {
			return err
		}

		r.Role = res.Role
	default:
		return ErrInvalidOrganizationRole
	}

	return nil
}

func (r *OrganizationRole) Value() (driver.Value, error) {
	if r == nil {
		return nil, nil
	}

	return r.Role, nil
}

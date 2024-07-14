package valueobjects

import (
	"database/sql/driver"
	"tenant-onboarding/modules/onboarding/internal/errorx"
)

var ErrInvalidTenantStatus = errorx.ErrInvalidTenantStatus

type TenantStatus struct {
	Status string
}

var (
	TenantCreated    = TenantStatus{"created"}
	TenantOnboarding = TenantStatus{"onboarding"}
	TenantActive     = TenantStatus{"activated"}
	TenantInactive   = TenantStatus{"inactive"}
)

func (s TenantStatus) String() string {
	return s.Status
}

func NewTenantStatus(s string) (TenantStatus, error) {
	switch s {
	case TenantCreated.Status:
		return TenantCreated, nil
	case TenantOnboarding.Status:
		return TenantOnboarding, nil
	case TenantActive.Status:
		return TenantActive, nil
	case TenantInactive.Status:
		return TenantInactive, nil
	}

	return TenantStatus{}, ErrInvalidTenantStatus
}

func (t *TenantStatus) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		res, err := NewTenantStatus(v)
		if err != nil {
			return err
		}

		t.Status = res.Status
	default:
		return ErrInvalidTenantStatus
	}

	return nil
}

func (t *TenantStatus) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}

	return t.String(), nil
}

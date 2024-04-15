package entity

import (
	"database/sql/driver"
	"errors"
)

type TenantStatus struct {
	status string
}

var (
	TenantCreated    = TenantStatus{"created"}
	TenantOnboarding = TenantStatus{"onboarding"}
	TenantActive     = TenantStatus{"active"}
	TenantInactive   = TenantStatus{"inactive"}
)

func (s TenantStatus) String() string {
	return s.status
}

func FromString(s string) (TenantStatus, error) {
	switch s {
	case TenantCreated.status:
		return TenantCreated, nil
	case TenantOnboarding.status:
		return TenantOnboarding, nil
	case TenantActive.status:
		return TenantActive, nil
	case TenantInactive.status:
		return TenantInactive, nil
	}

	return TenantStatus{}, errors.New("invalid status")
}

func (t *TenantStatus) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		res, err := FromString(v)
		if err != nil {
			return err
		}

		t.status = res.status
	default:
		return errors.New("unknown value")
	}

	return nil
}

func (t *TenantStatus) Value() (driver.Value, error) {
	if t == nil {
		return nil, nil
	}

	return t.String(), nil
}

package types

import "time"

type TenantOnboardedEvent struct {
	TenantID            string    `json:"tenant_id"`
	OrgID               string    `json:"org_id"`
	ProductID           string    `json:"product_id"`
	AppID               int       `json:"app_id"`
	Metadata            any       `json:"metadata"`
	ResourceInformation any       `json:"resource_information"`
	Timestamp           time.Time `json:"timestamp"`
}

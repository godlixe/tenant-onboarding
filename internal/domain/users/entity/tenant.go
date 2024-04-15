package entity

import (
	"tenant-onboarding/internal/domain/products"
	"tenant-onboarding/pkg/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	ID                  uuid.UUID           `json:"id"`
	ProductID           uuid.UUID           `json:"product_id" gorm:"column:product_id"`
	Product             products.Product    `json:"product" gorm:"foreignKey:product_id"`
	Name                string              `json:"name"`
	Subdomain           string              `json:"subdomain"`
	Status              TenantStatus        `json:"status"`
	ResourceInformation ResourceInformation `json:"resource_information"`
	database.Timestamp
}

func (t *Tenant) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.New()
	return nil
}

func NewTenant(
	ProductID uuid.UUID,
	Name string,
	Subdomain string,
) *Tenant {
	return &Tenant{
		ID:        uuid.New(),
		ProductID: ProductID,
		Name:      Name,
		Subdomain: Subdomain,
		Status:    TenantCreated,
	}
}

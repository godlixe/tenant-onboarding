package products

import "tenant-onboarding/pkg/database"

type Tier struct {
	ID    int
	Name  string
	Price int
	database.Timestamp
}

package products

import "tenant-onboarding/pkg/database"

type App struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	database.Timestamp
}

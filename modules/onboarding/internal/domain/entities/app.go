package entities

import (
	vo "tenant-onboarding/modules/onboarding/internal/domain/valueobjects"
	"tenant-onboarding/pkg/database"
	"tenant-onboarding/pkg/events/domain"
)

type App struct {
	ID   vo.AppID
	Name string

	events []domain.Event
	database.Timestamp
}

func NewApp(
	id vo.AppID,
	name string,
) *App {
	return &App{
		ID:   id,
		Name: name,
	}
}

func (a *App) Events() []domain.Event {
	return a.events
}

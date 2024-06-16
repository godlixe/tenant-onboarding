package providers

import "time"

func MarkAsFailed(app *App, event_name, listener_name, message string, metadata []byte, max_retries int) {
	app.DB.Table("failed_jobs").Create(map[string]any{
		"event_name":    event_name,
		"listener_name": listener_name,
		"metadata":      metadata,
		"max_retries":   max_retries,
		"created_at":    time.Now(),
		"message":       message,
	})
}

package framework

import "os"

// Checks if app is in integrated mode
func CheckIntegratedMode() bool {
	return os.Getenv("INTEGRATED_MODE") == "true"
}

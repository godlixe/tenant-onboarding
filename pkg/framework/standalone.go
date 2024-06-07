package framework

import "os"

// Checks if app is in integrated mode
func CheckIntegratedMode(any) bool {
	return os.Getenv("INTEGRATED_MODE") == "true"
}

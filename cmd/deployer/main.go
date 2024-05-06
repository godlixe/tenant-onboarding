package main

import (
	"fmt"
	"os"
	"tenant-onboarding/internal/worker"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	worker.Run()
}

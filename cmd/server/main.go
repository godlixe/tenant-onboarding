package main

import (
	"fmt"
	"os"
	"tenant-onboarding/middlewares"
	"tenant-onboarding/modules/auth"
	"tenant-onboarding/modules/onboarding"
	"tenant-onboarding/providers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := providers.NewApp()
	app.Webserver.Use(middlewares.CORSMiddleware())

	auth.RegisterModule(app)
	onboarding.RegisterModule(app)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	app.Webserver.Run(":" + port)

}

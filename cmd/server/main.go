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

	// creds, err := google.FindDefaultCredentials(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// client, err := pubsub.NewClient(
	// 	context.Background(),
	// 	os.Getenv("GOOGLE_PROJECT_ID"),
	// 	option.WithCredentialsJSON(creds.JSON),
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// db, err := database.NewPostgresClient(
	// 	database.DatabaseCredentials{
	// 		Host:     os.Getenv("DB_HOST"),
	// 		DBName:   os.Getenv("DB_NAME"),
	// 		User:     os.Getenv("DB_USER"),
	// 		Password: os.Getenv("DB_PASS"),
	// 		Port:     os.Getenv("DB_PORT"),
	// 	},
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// tenantRepository := postgresql.NewTenantRepository(db)
	// userRepository := postgresql.NewUserRepository(db)
	// tenantPublisher := queue.NewTenantPublisher(client)

	// tenantService := services.NewTenantService(tenantRepository, tenantPublisher)
	// authService := services.NewAuthService(userRepository)

	// tenantController := controllers.NewTenantController(tenantService)
	// authController := controllers.NewAuthController(authService)

	// routes.AuthRoutes(server, authController)
	// routes.TenantRoutes(server, tenantController)

	auth.RegisterModule(app)
	onboarding.RegisterModule(app)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	app.Webserver.Run(":" + port)

}

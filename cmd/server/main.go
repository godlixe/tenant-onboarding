package main

import (
	"context"
	"fmt"
	"os"
	"tenant-onboarding/internal/app/services"
	"tenant-onboarding/internal/infrastructures/database/postgresql"
	"tenant-onboarding/internal/infrastructures/queue/queue"
	"tenant-onboarding/internal/presentation/controllers"
	"tenant-onboarding/internal/presentation/middlewares"
	"tenant-onboarding/internal/presentation/routes"
	"tenant-onboarding/pkg/database"

	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	creds, err := google.FindDefaultCredentials(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client, err := pubsub.NewClient(
		context.Background(),
		os.Getenv("GOOGLE_PROJECT_ID"),
		option.WithCredentialsJSON(creds.JSON),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	db, err := database.NewPostgresClient(
		database.DatabaseCredentials{
			Host:     os.Getenv("DB_HOST"),
			DBName:   os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			Port:     os.Getenv("DB_PORT"),
		},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger := zerolog.New(os.Stdout)

	tenantRepository := postgresql.NewTenantRepository(db)
	userRepository := postgresql.NewUserRepository(db)
	tenantPublisher := queue.NewTenantPublisher(client)

	tenantService := services.NewTenantService(tenantRepository, tenantPublisher)
	authService := services.NewAuthService(userRepository)

	tenantController := controllers.NewTenantController(tenantService)
	authController := controllers.NewAuthController(authService)

	server := gin.Default()
	server.Use(middlewares.ErrorHandler(logger))

	routes.AuthRoutes(server, authController)
	routes.TenantRoutes(server, tenantController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}

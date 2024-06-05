package providers

import (
	"os"
	"tenant-onboarding/config"
	"tenant-onboarding/middlewares"

	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/samber/do"
	"gorm.io/gorm"
)

type App struct {
	DB        *gorm.DB
	Webserver *gin.Engine
	Logger    zerolog.Logger
	Queue     *pubsub.Client
	Injector  *do.Injector
}

func NewApp() *App {
	db := config.SetupDatabase(&config.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Name:     os.Getenv("DB_NAME"),
	})

	queue := config.InitPubsub()

	logger := zerolog.New(os.Stdout)

	server := gin.Default()

	server.Use(middlewares.ErrorHandler(logger))

	injector := do.New()

	return &App{
		DB:        db,
		Webserver: server,
		Logger:    logger,
		Queue:     queue,
		Injector:  injector,
	}
}

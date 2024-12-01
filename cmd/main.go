package main

import (
	"github.com/kossadda/wallet-service/models"
	"os"

	"github.com/joho/godotenv"
	"github.com/kossadda/wallet-service/pkg/handler"
	"github.com/kossadda/wallet-service/pkg/repository"
	"github.com/kossadda/wallet-service/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// @title Wallet Service API
// @version 1.0
// @description This is a wallet service API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := godotenv.Load("configs/config.env"); err != nil {
		logrus.Fatalf("error loading .env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.New(db)
	services := service.New(repos)
	handlers := handler.New(services)

	srv := models.NewServer()
	router := handlers.InitRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := srv.Run(port, router); err != nil {
		logrus.Fatalf("error while running http server: %s", err.Error())
	}
}

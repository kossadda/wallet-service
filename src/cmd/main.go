package main

import (
	"log"

	models "github.com/kossadda/wallet-service"
	"github.com/kossadda/wallet-service/pkg/handler"
	"github.com/kossadda/wallet-service/pkg/repository"
	"github.com/kossadda/wallet-service/pkg/service"
)

func main() {
	repos := repository.New()
	services := service.New(repos)
	handlers := handler.New(services)

	srv := models.NewServer()
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}

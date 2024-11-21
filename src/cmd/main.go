package main

import (
	"log"

	models "github.com/kossadda/wallet-service"
	"github.com/kossadda/wallet-service/pkg/handler"
)

func main() {
	handlers := handler.New()
	srv := models.NewServer()

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}

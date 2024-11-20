package main

import (
	"log"

	models "github.com/kossadda/wallet-service"
)

func main() {
	srv := new(models.Server)

	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}

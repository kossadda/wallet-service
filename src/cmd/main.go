package main

import (
	"log"

	hs "github.com/kossadda/wallet-service"
)

func main() {
	srv := new(hs.Server)

	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error while running http server: %s", err.Error())
	}
}

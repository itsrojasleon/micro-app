package main

import (
	"log"

	"github.com/rojasleon/reserve-micro/auth/models"
)

func main() {
	ValidateEnvVars()
	models.ConnectToDatabase()

	r := SetupRouter()

	log.Fatal(r.Run(":8000"))
}

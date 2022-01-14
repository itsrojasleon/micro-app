package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rojasleon/reserve-micro/auth/routes"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	routes.InitAuthRouter(r)

	return r
}

// Make sure needed environment variables are defined
func ValidateEnvVars() {
	keys := []string{"JWT_SECRET"}

	for _, key := range keys {
		if os.Getenv(key) == "" {
			log.Fatal(key + " must be defined")
		}
	}
}

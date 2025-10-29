package main

import (
	"cli-arithmetic-app/rest-api/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadENV()
	host := os.Getenv("REST_API_HOST")
	port := os.Getenv("REST_API_PORT")
	address := fmt.Sprintf("%s:%s", host, port)

	r := gin.Default()

	routes.RegisterRoutes(r)

	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	} else {
		log.Printf("REST API server running on http://%s:%s", host, port)
	}
}

func loadENV() {
	if err := godotenv.Load(); err != nil {
		// logger.Log.Warn("No .env file found (skipping)")
		fmt.Print("No .env file found (skipping)")
	}
}

package main

import (
	"cli-arithmetic-app/cli"
	logger "cli-arithmetic-app/log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	defer handlePanic()

	logger.InitLogger()

	logger.Log.Info("===========================================================================================")
	logger.Log.Info("App inicialization...")

	// Add environment variables
	if err := godotenv.Load(); err != nil {
		logger.Log.Warn("No .env file found (skipping)")
	}

	cli.Execute()
}

func handlePanic() {
	if err := recover(); err != nil {
		logger.Log.Fatalf("Error: %v", err)
		os.Exit(1)
	}
}

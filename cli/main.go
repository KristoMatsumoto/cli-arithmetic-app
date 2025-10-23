package main

import (
	"fmt"
	"os"
)

func main() {
	defer handlePanic()

	// logger.InitLogger()

	// logger.Log.Info("===========================================================================================")
	// logger.Log.Info("App inicialization...")
	fmt.Print("===========================================================================================")
	fmt.Print("App inicialization...")

	// Add environment variables
	// if err := godotenv.Load(); err != nil {
	// 	// logger.Log.Warn("No .env file found (skipping)")
	// 	fmt.Print("No .env file found (skipping)")
	// }

	Execute()
}

func handlePanic() {
	if err := recover(); err != nil {
		// logger.Log.Fatalf("Error: %v", err)
		fmt.Errorf("Error: %v", err)
		os.Exit(1)
	}
}

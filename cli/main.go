package main

import (
	"cli-arithmetic-app/cli/core"
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

func Execute() {
	if err := core.RootCmd.Execute(); err != nil {
		// logger.Log.Fatalf("Error starting CLI: %v", err)
		fmt.Printf("Error starting CLI: %v", err)
		os.Exit(1)
	}
}

func handlePanic() {
	if err := recover(); err != nil {
		// logger.Log.Fatalf("Error: %v", err)
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

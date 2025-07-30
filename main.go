package main

import (
	"cli-arithmetic-app/cli"
	logger "cli-arithmetic-app/log"
	"os"
)

func main() {
	defer handlePanic()

	logger.InitLogger()

	logger.Log.Info("===========================================================================================")
	logger.Log.Info("App inicialization...")
	cli.Execute()
}

func handlePanic() {
	if err := recover(); err != nil {
		logger.Log.Fatalf("Error: %v", err)
		os.Exit(1)
	}
}

package cli

import (
	"cli-arithmetic-app/app"
	logger "cli-arithmetic-app/log"
	"os"

	"github.com/spf13/cobra"
)

var (
	version   string
	processor string
	inPath    string
	outPath   string
	format    string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&version, "version", "v", "1", "Logic version (1 or 2)")
	rootCmd.PersistentFlags().StringVarP(&processor, "processor-version", "p", "1", "Processor version (1, 2 or 3)")
	rootCmd.PersistentFlags().StringVarP(&inPath, "in", "i", "test/inputs/input.txt", "Input file path")
	rootCmd.PersistentFlags().StringVarP(&outPath, "out", "o", "test/outputs/output", "Output file path")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "txt", "Input/output format (txt, html, json, xml, yaml)")
}

var rootCmd = &cobra.Command{
	Use:   "cli-arithmetic-app",
	Short: "Cli application for processing files with arithmetic expressions",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Log.Fatalf("Error starting CLI: %v", err)
		os.Exit(1)
	}
}

func run() {
	opts := app.PipelineOptions{
		InputPath:     inPath,
		OutputPath:    outPath + "." + format,
		Format:        format,
		ProcessorType: processor,
		Version:       version,
	}

	if err := app.ExecuteProcessingPipeline(opts); err != nil {
		logger.Log.Fatalf("Pipeline failed: %v", err)
	}
}

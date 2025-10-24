package core

import (
	"cli-arithmetic-app/app/core"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version        string
	processor      string
	inPath         string
	outPath        string
	format         string
	transformChain []string
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&version, "version", "v", "1", "Logic version (1 or 2)")
	RootCmd.PersistentFlags().StringVarP(&processor, "processor", "p", "1", "Processor to use (1, 2 or 3)")
	RootCmd.PersistentFlags().StringVarP(&inPath, "in", "i", "app/test/inputs/input.txt", "Input file path")
	RootCmd.PersistentFlags().StringVarP(&outPath, "out", "o", "app/test/outputs/output", "Output file path")
	RootCmd.PersistentFlags().StringVarP(&format, "format", "f", "txt", "Input/output format (txt, html, json, xml, yaml)")
	RootCmd.PersistentFlags().StringSliceVarP(&transformChain, "transform", "t", []string{}, "Comma-separated list of transformers (zip, aes, etc.)")
}

var RootCmd = &cobra.Command{
	Use:   "cli-arithmetic-app",
	Short: "Cli application for processing files with arithmetic expressions",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	opts := core.PipelineOptions{
		InputPath:      inPath,
		OutputPath:     outPath + "." + format,
		Format:         format,
		ProcessorType:  processor,
		Version:        version,
		TransformChain: transformChain,
	}

	if err := core.ExecuteProcessingPipeline(opts); err != nil {
		// logger.Log.Fatalf("Pipeline failed: %v", err)
		fmt.Printf("Pipeline failed: %v", err)
	}
}

package app

import (
	logger "cli-arithmetic-app/log"
	"cli-arithmetic-app/modules/v1/parser"
	"cli-arithmetic-app/modules/v1/processor"
)

func ExecuteProcessingPipeline(options PipelineOptions) error {
	logger.Log.Infof("Executing logic version %s", options.Version)
	logger.Log.Infof("Running with input: %s, output: %s, format: %s", options.InputPath, options.OutputPath, options.Format)

	// TODO:
	// 		Change format parameter for input
	// 		Add encryption
	// 		Add archiving

	// Get parser for input
	// var format string
	// if strings.Contains(options.InputPath, ".") {
	// 	format = strings.TrimPrefix(filepath.Ext(options.InputPath), ".")
	// } else {
	// 	return fmt.Errorf("cannot read input file format from %s", options.InputPath)
	// }

	// Parsing
	parserInstance, err := parser.CreateParser(options.Format)
	if err != nil {
		return err
	}

	// Read file
	lines, err := parserInstance.ReadFile(options.InputPath)
	if err != nil {
		return err
	}
	logger.Log.Infof("Read %d lines", len(lines))

	// Get processor
	logger.Log.Infof("Starting processor %s", options.ProcessorType)
	processorInstance, err := processor.CreateProcessor(options.ProcessorType)
	if err != nil {
		return err
	}

	// Process
	processed, err := processorInstance.Process(lines)
	if err != nil {
		return err
	}

	// Write output
	err = parserInstance.WriteFile(options.OutputPath, processed)
	if err != nil {
		return err
	}

	logger.Log.Info("Processing pipeline finished successfully")
	return nil
}

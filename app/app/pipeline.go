package app

import (
	logger "cli-arithmetic-app/log"
	"fmt"
	"os"
)

func ExecuteProcessingPipeline(options PipelineOptions) error {
	logger.Log.Infof("Executing logic version %s", options.Version)
	logger.Log.Infof(
		"Running with input: %s, output: %s, format: %s, processor: %s, transforms: %v",
		options.InputPath, options.OutputPath, options.Format, options.ProcessorType, options.TransformChain)

	// Decode
	chain, err := BuildTransformChain(options.TransformChain)
	if err != nil {
		return err
	}
	logger.Log.Infof("File transforms chain (%v) has been added", options.TransformChain)

	fileData, err := os.ReadFile(options.InputPath)
	if err != nil {
		return err
	}
	logger.Log.Info("File data has been read")

	for _, t := range chain {
		logger.Log.Infof("Decoding with %s", t.Name())
		fileData, err = t.Decode(fileData)
		if err != nil {
			return fmt.Errorf("decode (%s): %w", t.Name(), err)
		}
	}

	// Parsing
	parserInstance, err := CreateParser(options.Format)
	if err != nil {
		return err
	}
	logger.Log.Info("The parser has been selected")

	// Read file
	lines, err := parserInstance.ParseBytes(fileData)
	if err != nil {
		return err
	}
	logger.Log.Infof("Read %d lines", len(lines))

	// Get processor
	logger.Log.Infof("Starting processor %s", options.ProcessorType)
	processorInstance, err := CreateProcessor(options.ProcessorType)
	if err != nil {
		return err
	}

	// Process
	processed, err := processorInstance.Process(lines)
	if err != nil {
		return err
	}
	logger.Log.Infof("Processing completed: %d lines", len(processed))

	// Write output
	outBytes, err := parserInstance.SerializeBytes(processed)
	if err != nil {
		return err
	}
	logger.Log.Info("Data has been returned to bytes")

	// data := []byte(strings.Join(processed, "\n"))

	// Encode & Write
	for _, t := range chain {
		logger.Log.Infof("Encoding with %s", t.Name())
		outBytes, err = t.Encode(outBytes)
		if err != nil {
			return fmt.Errorf("encode (%s): %w", t.Name(), err)
		}
	}

	err = os.WriteFile(options.OutputPath, outBytes, 0644)
	if err != nil {
		return err
	}

	logger.Log.Info("Processing pipeline finished successfully")
	return nil
}

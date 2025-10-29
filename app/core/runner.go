package core

import (
	logger "cli-arithmetic-app/app/log"
)

// =========================================================
//                         DECORATORS
// =========================================================

func WithLogger(f func() error) error {
	logger.InitLogger()
	return f()
}

// =========================================================
//                         RUNNERS
// =========================================================

func ExecuteProcessingPipeline(options PipelineOptions) error {
	return WithLogger(func() error {
		return pipeline(options)
	})
}

func ExecuteProcess(lines []string, processorType string) ([]string, error) {
	return process(lines, processorType)
}

func ExecuteParse(bytes []byte, format string) ([]string, error) {
	return parse(bytes, format)
}

func ExecuteCompose(lines []string, format string) ([]byte, error) {
	return compose(lines, format)
}

func ExecuteTransform(bytes []byte, format string) ([]byte, error) {
	return transform(bytes, format)
}

func ExecuteTransforms(bytes []byte, formats []string) ([]byte, error) {
	return transformWithChain(bytes, formats)
}

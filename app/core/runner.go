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

func ExecuteEncode(bytes []byte, format string) ([]byte, error) {
	return encode(bytes, format)
}

func ExecuteEncodeMany(bytes []byte, formats []string) ([]byte, error) {
	return encodeWithChain(bytes, formats)
}

func ExecuteDecode(bytes []byte, format string) ([]byte, error) {
	return decode(bytes, format)
}

func ExecuteDecodeMany(bytes []byte, formats []string) ([]byte, error) {
	return decodeWithChain(bytes, formats)
}

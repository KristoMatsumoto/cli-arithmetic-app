package processor

import (
	"cli-arithmetic-app/modules/v1/processor/naive_processor"
	"fmt"
)

// Processor defines interface for arithmetic expression processing
type Processor interface {
	Process(lines []string) ([]string, error)
}

// CreateProcessor returns processor based on implementation type
func CreateProcessor(mode string) (Processor, error) {
	switch mode {
	case "1":
		return naive_processor.NewNaiveProcessor(), nil
	case "2":
		return NewRegexProcessor(), nil
	default:
		return nil, fmt.Errorf("unsupported processor mode: %s", mode)
	}
}

package processor

// Processor defines interface for arithmetic expression processing
type Processor interface {
	Process(lines []string) ([]string, error)
}

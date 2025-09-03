package parser

// Parser defines a common interface for all file parsers (text, json, xml, etc.)
type Parser interface {
	ReadFile(path string) ([]string, error)     // Parse input into raw lines or expressions
	WriteFile(path string, data []string) error // Save processed output
}

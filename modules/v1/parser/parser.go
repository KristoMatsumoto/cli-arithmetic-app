package parser

import "fmt"

// Parser defines a common interface for all file parsers (text, json, xml, etc.)
type Parser interface {
	ReadFile(path string) ([]string, error)     // Parse input into raw lines or expressions
	WriteFile(path string, data []string) error // Save processed output
}

func CreateParser(format string) (Parser, error) {
	switch format {
	case "text":
		return NewTextParser(), nil
	case "txt":
		return NewTextParser(), nil
	case "json":
		return NewJSONParser(), nil
	case "xml":
		return NewXMLParser(), nil
	case "yaml":
		return NewYAMLParser(), nil
	case "html":
		return NewHTMLParser(), nil
	default:
		return nil, fmt.Errorf("unsupported format")
	}
}

package parser

// Parser defines a common interface for all file parsers (text, json, xml, etc.)
type Parser interface {
	ReadFile(path string) ([]string, error)
	ParseBytes(raw []byte) ([]string, error)
	WriteFile(path string, data []string) error
	SerializeBytes(data []string) ([]byte, error)
}

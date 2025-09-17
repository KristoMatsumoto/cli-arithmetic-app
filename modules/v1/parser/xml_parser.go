package parser

import (
	"encoding/xml"
	"os"
)

type XMLParser struct{}

func NewXMLParser() *XMLParser {
	return &XMLParser{}
}

type xmlWrapper struct {
	Lines []string `xml:"line"`
}

func (parser *XMLParser) ReadFile(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parser.ParseBytes(file)
}

func (parser *XMLParser) ParseBytes(raw []byte) ([]string, error) {
	var wrapper xmlWrapper
	if err := xml.Unmarshal(raw, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.Lines, nil
}

func (parser *XMLParser) WriteFile(path string, data []string) error {
	raw, err := parser.SerializeBytes(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0644)
}

func (parser *XMLParser) SerializeBytes(data []string) ([]byte, error) {
	wrapper := xmlWrapper{Lines: data}
	return xml.MarshalIndent(wrapper, "", "  ")
}

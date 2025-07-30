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

func (x *XMLParser) ReadFile(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wrapper xmlWrapper
	if err := xml.Unmarshal(file, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.Lines, nil
}

func (x *XMLParser) WriteFile(path string, data []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	wrapper := xmlWrapper{Lines: data}
	output, err := xml.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(output)
	return err
}

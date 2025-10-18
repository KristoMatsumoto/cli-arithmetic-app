package parser

import (
	"encoding/json"
	"os"
)

type JSONParser struct{}

func NewJSONParser() *JSONParser {
	return &JSONParser{}
}

func (parser *JSONParser) Format() string {
	return "json"
}

func (parser *JSONParser) ReadFile(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return parser.ParseBytes(file)
}

func (parser *JSONParser) ParseBytes(raw []byte) ([]string, error) {
	var data []string
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (parser *JSONParser) WriteFile(path string, data []string) error {
	raw, err := parser.SerializeBytes(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0644)
}

func (parser *JSONParser) SerializeBytes(data []string) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}

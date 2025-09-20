package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

type YAMLParser struct{}

func NewYAMLParser() *YAMLParser {
	return &YAMLParser{}
}

func (parser *YAMLParser) ReadFile(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parser.ParseBytes(file)
}

func (parser *YAMLParser) ParseBytes(raw []byte) ([]string, error) {
	var data []string
	if err := yaml.Unmarshal(raw, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (parser *YAMLParser) WriteFile(path string, data []string) error {
	raw, err := parser.SerializeBytes(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0644)
}

func (parser *YAMLParser) SerializeBytes(data []string) ([]byte, error) {
	return yaml.Marshal(data)
}

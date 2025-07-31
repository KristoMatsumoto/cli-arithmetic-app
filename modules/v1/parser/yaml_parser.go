package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type YAMLParser struct{}

func NewYAMLParser() *YAMLParser {
	return &YAMLParser{}
}

func (y *YAMLParser) ReadFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	if err := yaml.Unmarshal(data, &lines); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return lines, nil
}

func (y *YAMLParser) WriteFile(path string, data []string) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(path, yamlData, 0644)
}

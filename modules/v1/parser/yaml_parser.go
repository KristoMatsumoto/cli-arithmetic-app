package parser

import (
	"os"

	"gopkg.in/yaml.v3"
)

type YAMLParser struct{}

func NewYAMLParser() *YAMLParser {
	return &YAMLParser{}
}

type yamlWrapper struct {
	Lines []string `yaml:"lines"`
}

func (y *YAMLParser) ReadFile(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wrapper yamlWrapper
	if err := yaml.Unmarshal(file, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.Lines, nil
}

func (y *YAMLParser) WriteFile(path string, data []string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	wrapper := yamlWrapper{Lines: data}
	output, err := yaml.Marshal(wrapper)
	if err != nil {
		return err
	}

	_, err = file.Write(output)
	return err
}

package parser

import (
	"bytes"
	"os"
	"strings"
)

type TextParser struct{}

func NewTextParser() *TextParser {
	return &TextParser{}
}

func (parser *TextParser) Format() string {
	return "txt"
}

func (parser *TextParser) ReadFile(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parser.ParseBytes(file)
}

func (parser *TextParser) ParseBytes(raw []byte) ([]string, error) {
	content := string(raw)
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
	return lines, nil
}

func (parser *TextParser) WriteFile(path string, data []string) error {
	raw, err := parser.SerializeBytes(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0644)
}

func (parser *TextParser) SerializeBytes(data []string) ([]byte, error) {
	var buf bytes.Buffer
	for i, line := range data {
		buf.WriteString(line)
		if i < len(data)-1 {
			buf.WriteString("\n")
		}
	}
	return buf.Bytes(), nil
}

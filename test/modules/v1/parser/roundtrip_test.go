package parser_test

import (
	"os"
	"path/filepath"
	"testing"

	"cli-arithmetic-app/modules/v1/parser"

	"github.com/stretchr/testify/assert"
)

func roundtrip(t *testing.T, p parser.Parser, inputPath string) {
	tempFile := filepath.Join(os.TempDir(), "roundtrip_output.tmp")

	original, err := p.ReadFile(inputPath)
	assert.NoError(t, err, "reading original file")

	err = p.WriteFile(tempFile, original)
	assert.NoError(t, err, "writing roundtrip file")

	reconstructed, err := p.ReadFile(tempFile)
	assert.NoError(t, err, "reading roundtrip file")

	assert.Equal(t, original, reconstructed, "roundtrip mismatch")
}

func TestRoundtripHTML(t *testing.T) {
	p := parser.NewHTMLParser()
	wd, _ := os.Getwd()
	inputPath := filepath.Join(wd, "html_parser", "sample.html")
	roundtrip(t, p, inputPath)
}

func TestRoundtripJSON(t *testing.T) {
	p := parser.NewJSONParser()
	wd, _ := os.Getwd()
	inputPath := filepath.Join(wd, "json_parser", "sample.json")
	roundtrip(t, p, inputPath)
}

func TestRoundtripTXT(t *testing.T) {
	p := parser.NewTextParser()
	wd, _ := os.Getwd()
	inputPath := filepath.Join(wd, "text_parser", "sample.txt")
	roundtrip(t, p, inputPath)
}

func TestRoundtripXML(t *testing.T) {
	p := parser.NewXMLParser()
	wd, _ := os.Getwd()
	inputPath := filepath.Join(wd, "xml_parser", "sample.xml")
	// inputPath := "test/modules/v1/parser/xml_parser/sample.xml"
	roundtrip(t, p, inputPath)
}

func TestRoundtripYAML(t *testing.T) {
	p := parser.NewYAMLParser()
	wd, _ := os.Getwd()
	inputPath := filepath.Join(wd, "yaml_parser", "sample.yaml")
	roundtrip(t, p, inputPath)
}

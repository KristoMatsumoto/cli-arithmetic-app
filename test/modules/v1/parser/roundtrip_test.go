package parser_test

import (
	"os"
	"path/filepath"
	"testing"

	"cli-arithmetic-app/modules/v1/parser"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/stretchr/testify/assert"
)

func roundtrip(t provider.T, p parser.Parser, inputPath string) {
	tempFile := filepath.Join(os.TempDir(), "roundtrip_output.tmp")

	var original []string
	t.WithNewStep("Read original file", func(sCtx provider.StepCtx) {
		var err error
		original, err = p.ReadFile(inputPath)
		assert.NoError(t, err, "reading original file")
	})

	t.WithNewStep("Write content to temp file", func(sCtx provider.StepCtx) {
		err := p.WriteFile(tempFile, original)
		assert.NoError(t, err, "writing roundtrip file")
	})

	var reconstructed []string
	t.WithNewStep("Read back from temp file", func(sCtx provider.StepCtx) {
		var err error
		reconstructed, err = p.ReadFile(tempFile)
		assert.NoError(t, err, "reading roundtrip file")
	})

	t.WithNewStep("Compare original and reconstructed content", func(sCtx provider.StepCtx) {
		assert.Equal(t, original, reconstructed, "roundtrip mismatch")
	})

	_ = os.Remove(tempFile)
}

func TestHTMLParser_Roundtrip(t *testing.T) {
	runner.Run(t, "Roundtrip HTML Parser", func(t provider.T) {
		p := parser.NewHTMLParser()
		wd, _ := os.Getwd()
		inputPath := filepath.Join(wd, "html_parser", "sample.html")
		roundtrip(t, p, inputPath)
	})
}

func TestJSONParser_Roundtrip(t *testing.T) {
	runner.Run(t, "Roundtrip JSON Parser", func(t provider.T) {
		p := parser.NewJSONParser()
		wd, _ := os.Getwd()
		inputPath := filepath.Join(wd, "json_parser", "sample.json")
		roundtrip(t, p, inputPath)
	})
}

func TestTXTParser_Roundtrip(t *testing.T) {
	runner.Run(t, "Roundtrip Text Parser", func(t provider.T) {
		p := parser.NewTextParser()
		wd, _ := os.Getwd()
		inputPath := filepath.Join(wd, "text_parser", "sample.txt")
		roundtrip(t, p, inputPath)
	})
}

func TestXMLParser_Roundtrip(t *testing.T) {
	runner.Run(t, "Roundtrip XML Parser", func(t provider.T) {
		p := parser.NewXMLParser()
		wd, _ := os.Getwd()
		inputPath := filepath.Join(wd, "xml_parser", "sample.xml")
		roundtrip(t, p, inputPath)
	})
}

func TestYAMLParser_Roundtrip(t *testing.T) {
	runner.Run(t, "Roundtrip YAML Parser", func(t provider.T) {
		p := parser.NewYAMLParser()
		wd, _ := os.Getwd()
		inputPath := filepath.Join(wd, "yaml_parser", "sample.yaml")
		roundtrip(t, p, inputPath)
	})
}

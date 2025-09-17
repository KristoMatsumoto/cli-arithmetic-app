package yaml_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"cli-arithmetic-app/utils/parsertest"
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

var input = []string{"10/2", "6*7", "World"}

func TestYAMLParser_ReadWrite(t *testing.T) {
	runner.Run(t, "YAML Parser: Read & Write", func(t provider.T) {
		p := parser.NewYAMLParser()
		tempFile := filepath.Join(os.TempDir(), "yaml_test.yaml")

		t.WithNewStep("Write data to temp YAML file", func(sCtx provider.StepCtx) {
			err := p.WriteFile(tempFile, input)
			t.Assert().NoError(err)
		})

		var output []string
		t.WithNewStep("Read data back from temp YAML file", func(sCtx provider.StepCtx) {
			var err error
			output, err = p.ReadFile(tempFile)
			t.Assert().NoError(err)
		})

		t.WithNewStep("Compare written and read data", func(sCtx provider.StepCtx) {
			t.Assert().Equal(input, output)
		})

		_ = os.Remove(tempFile)
	})
}

func TestYAMLParser_Roundtrip(t *testing.T) {
	runner.Run(t, "Roundtrip YAML Parser", func(t provider.T) {
		p := parser.NewYAMLParser()
		wd, _ := os.Getwd()
		inputPath := filepath.Join(wd, "sample.yaml")
		parsertest.Roundtrip(t, p, inputPath)
	})
}

func TestYAMLParser_RoundtripBytes(t *testing.T) {
	runner.Run(t, "Roundtrip bytes YAML Parser", func(t provider.T) {
		p := parser.NewYAMLParser()
		parsertest.RoundtripBytes(t, p, input)
	})
}

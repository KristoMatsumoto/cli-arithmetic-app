package json_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"cli-arithmetic-app/utils/roundtrip"
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

var input = []string{"4+4", "5-2", "Test"}

func TestJSONParser_ReadWrite(t *testing.T) {
	runner.Run(t, "JSON Parser: Read & Write", func(t provider.T) {
		p := parser.NewJSONParser()
		tempFile := filepath.Join(os.TempDir(), "json_test.json")

		t.WithNewStep("Write data to temp JSON file", func(sCtx provider.StepCtx) {
			err := p.WriteFile(tempFile, input)
			t.Assert().NoError(err)
		})

		var output []string
		t.WithNewStep("Read data back from temp JSON file", func(sCtx provider.StepCtx) {
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

func TestJSONParser_Roundtrip(t *testing.T) {
	runner.Run(t, "Roundtrip JSON Parser", func(t provider.T) {
		p := parser.NewJSONParser()
		wd, _ := os.Getwd()
		inputPath := filepath.Join(wd, "sample.json")
		roundtrip.Roundtrip(t, p, inputPath)
	})
}

func TestJSONParser_RoundtripBytes(t *testing.T) {
	runner.Run(t, "Roundtrip bytes JSON Parser", func(t provider.T) {
		p := parser.NewJSONParser()
		roundtrip.RoundtripBytes(t, p, input)
	})
}

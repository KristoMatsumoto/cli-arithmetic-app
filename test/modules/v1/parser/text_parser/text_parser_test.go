package text_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"cli-arithmetic-app/utils/roundtrip"
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

var input = []string{"1+1", "2*3", "Hello"}

func TestTXTParser_ReadWrite(t *testing.T) {
	runner.Run(t, "TXT Parser: Read & Write", func(t provider.T) {
		p := parser.NewTextParser()
		tempFile := filepath.Join(os.TempDir(), "text_test.txt")

		t.WithNewStep("Write data to temp TXT file", func(sCtx provider.StepCtx) {
			err := p.WriteFile(tempFile, input)
			t.Assert().NoError(err)
		})

		var output []string
		t.WithNewStep("Read data back from temp TXT file", func(sCtx provider.StepCtx) {
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

func TestTXTParser_Roundtrip(t *testing.T) {
	runner.Run(t, "Roundtrip Text Parser", func(t provider.T) {
		p := parser.NewTextParser()
		wd, _ := os.Getwd()
		inputPath := filepath.Join(wd, "sample.txt")
		roundtrip.Roundtrip(t, p, inputPath)
	})
}

func TestTXTParser_RoundtripBytes(t *testing.T) {
	runner.Run(t, "Roundtrip bytes Text Parser", func(t provider.T) {
		p := parser.NewTextParser()
		roundtrip.RoundtripBytes(t, p, input)
	})
}

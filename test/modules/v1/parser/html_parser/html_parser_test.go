package html_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"cli-arithmetic-app/utils/roundtrip"
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

var input = []string{"3+3", "<b>bold</b>", "simple"}

func TestHTMLParser_ReadWrite(t *testing.T) {
	runner.Run(t, "HTML Parser: Read & Write", func(t provider.T) {
		p := parser.NewHTMLParser()
		tempFile := filepath.Join(os.TempDir(), "html_test.html")

		t.WithNewStep("Write data to temp HTML file", func(sCtx provider.StepCtx) {
			err := p.WriteFile(tempFile, input)
			t.Assert().NoError(err)
		})

		var output []string
		t.WithNewStep("Read data back from temp HTML file", func(sCtx provider.StepCtx) {
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

func TestHTMLParser_Roundtrip(t *testing.T) {
	runner.Run(t, "Roundtrip HTML Parser", func(t provider.T) {
		p := parser.NewHTMLParser()
		wd, _ := os.Getwd()
		inputPath := filepath.Join(wd, "sample.html")
		roundtrip.Roundtrip(t, p, inputPath)
	})
}

func TestHTMLParser_RoundtripBytes(t *testing.T) {
	runner.Run(t, "Roundtrip bytes HTML Parser", func(t provider.T) {
		p := parser.NewHTMLParser()
		roundtrip.RoundtripBytes(t, p, input)
	})
}

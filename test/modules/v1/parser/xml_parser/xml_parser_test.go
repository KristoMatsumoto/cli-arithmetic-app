package xml_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/stretchr/testify/assert"
)

func TestXMLParser_ReadWrite(t *testing.T) {
	runner.Run(t, "XML Parser: Read & Write", func(t provider.T) {
		p := parser.NewXMLParser()
		tempFile := filepath.Join(os.TempDir(), "xml_test.xml")
		input := []string{"1+2", "8/4", "XMLLine"}

		t.WithNewStep("Write data to temp XML file", func(sCtx provider.StepCtx) {
			err := p.WriteFile(tempFile, input)
			assert.NoError(t, err)
		})

		var output []string
		t.WithNewStep("Read data back from temp XML file", func(sCtx provider.StepCtx) {
			var err error
			output, err = p.ReadFile(tempFile)
			assert.NoError(t, err)
		})

		t.WithNewStep("Compare written and read data", func(sCtx provider.StepCtx) {
			assert.Equal(t, input, output)
		})

		_ = os.Remove(tempFile)
	})
}

package yaml_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"os"
	"path/filepath"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/stretchr/testify/assert"
)

func TestYAMLParser_ReadWrite(t *testing.T) {
	runner.Run(t, "YAML Parser: Read & Write", func(t provider.T) {
		p := parser.NewYAMLParser()
		tempFile := filepath.Join(os.TempDir(), "yaml_test.yaml")
		input := []string{"10/2", "6*7", "World"}

		t.WithNewStep("Write data to temp YAML file", func(sCtx provider.StepCtx) {
			err := p.WriteFile(tempFile, input)
			assert.NoError(t, err)
		})

		var output []string
		t.WithNewStep("Read data back from temp YAML file", func(sCtx provider.StepCtx) {
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

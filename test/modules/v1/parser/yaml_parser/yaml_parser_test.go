package yaml_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYAMLParser_ReadWrite(t *testing.T) {
	p := parser.NewYAMLParser()
	tempFile := filepath.Join(os.TempDir(), "yaml_test.yaml")
	input := []string{"10/2", "6*7", "World"}

	err := p.WriteFile(tempFile, input)
	assert.NoError(t, err)

	output, err := p.ReadFile(tempFile)
	assert.NoError(t, err)
	assert.Equal(t, input, output)

	_ = os.Remove(tempFile)
}

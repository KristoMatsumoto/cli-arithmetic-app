package json_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONParser_ReadWrite(t *testing.T) {
	p := parser.NewJSONParser()
	tempFile := filepath.Join(os.TempDir(), "json_test.json")
	input := []string{"4+4", "5-2", "Test"}

	err := p.WriteFile(tempFile, input)
	assert.NoError(t, err)

	output, err := p.ReadFile(tempFile)
	assert.NoError(t, err)
	assert.Equal(t, input, output)

	_ = os.Remove(tempFile)
}

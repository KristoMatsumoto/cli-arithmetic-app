package text_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextParser_ReadWrite(t *testing.T) {
	p := parser.NewTextParser()
	tempFile := filepath.Join(os.TempDir(), "text_test.txt")
	input := []string{"1+1", "2*3", "Hello"}

	err := p.WriteFile(tempFile, input)
	assert.NoError(t, err)

	output, err := p.ReadFile(tempFile)
	assert.NoError(t, err)
	assert.Equal(t, input, output)

	_ = os.Remove(tempFile)
}

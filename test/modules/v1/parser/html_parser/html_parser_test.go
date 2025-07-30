package html_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTMLParser_ReadWrite(t *testing.T) {
	p := parser.NewHTMLParser()
	tempFile := filepath.Join(os.TempDir(), "html_test.html")
	input := []string{"3+3", "<b>bold</b>", "simple"}

	err := p.WriteFile(tempFile, input)
	assert.NoError(t, err)

	output, err := p.ReadFile(tempFile)
	assert.NoError(t, err)
	assert.Equal(t, input, output)

	_ = os.Remove(tempFile)
}

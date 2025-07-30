package xml_parser_test

import (
	"cli-arithmetic-app/modules/v1/parser"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXMLParser_ReadWrite(t *testing.T) {
	p := parser.NewXMLParser()
	tempFile := filepath.Join(os.TempDir(), "xml_test.xml")
	input := []string{"1+2", "8/4", "XMLLine"}

	err := p.WriteFile(tempFile, input)
	assert.NoError(t, err)

	output, err := p.ReadFile(tempFile)
	assert.NoError(t, err)
	assert.Equal(t, input, output)

	_ = os.Remove(tempFile)
}

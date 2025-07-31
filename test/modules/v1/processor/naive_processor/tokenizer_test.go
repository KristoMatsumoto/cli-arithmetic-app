package naive_processor_test

import (
	"testing"

	"cli-arithmetic-app/modules/v1/processor"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	input := "3 + 4 * (2 - 1)^2 % 5"
	tokens, err := processor.Tokenize(input)
	assert.NoError(t, err)

	expected := []processor.Token{
		{Type: processor.Number, Value: "3"},
		{Type: processor.Operator, Value: "+"},
		{Type: processor.Number, Value: "4"},
		{Type: processor.Operator, Value: "*"},
		{Type: processor.LParen, Value: "("},
		{Type: processor.Number, Value: "2"},
		{Type: processor.Operator, Value: "-"},
		{Type: processor.Number, Value: "1"},
		{Type: processor.RParen, Value: ")"},
		{Type: processor.Operator, Value: "^"},
		{Type: processor.Number, Value: "2"},
		{Type: processor.Operator, Value: "%"},
		{Type: processor.Number, Value: "5"},
	}

	assert.Equal(t, expected, tokens)
}

func TestTokenize_InvalidCharacter(t *testing.T) {
	_, err := processor.Tokenize("2 + x")
	assert.Error(t, err)
}
